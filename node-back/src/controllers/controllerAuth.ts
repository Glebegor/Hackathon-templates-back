import { PrismaClient } from "@prisma/client";
import { IConfig } from "../core/commonServer/config";
import { IRepositoryAuth, newRepositoryAuth } from "../repositories/repositoryAuth";
import { IUsecaseAuth, newUsecaseAuth } from "../usecases/usecaseAuth";
import { IRegisterRequest, ILoginRequest, IRefreshTokenRequest } from "../core/commonApi/authApi";
import { hashPassword } from "../utils/passwordHash";
import { ResponseSuccess } from "../core/commonApi/responseSuccess";
import { ResponseError } from "../core/commonApi/responseError";
import { IUser } from "../core/common/user";
import { generateAccessToken, generateRefreshToken, verifyToken } from "../utils/jwtTokens";
const bcrypt = require('bcrypt');

class ControllerAuth {
    public config: IConfig;
    public usecase: IUsecaseAuth;
    public routeString: string;

    constructor(config: IConfig, databaseClient: PrismaClient, routeString: string) {
        var repo: IRepositoryAuth = newRepositoryAuth(databaseClient);
        var usecase: IUsecaseAuth = newUsecaseAuth(repo);

        this.config = config;
        this.usecase = usecase;
        this.routeString = routeString;
    }

    private async register(req: any, res: any): Promise<void> {
        const input: IRegisterRequest = req.body;

        var hashedPassword: string = hashPassword(input.password, this.config.SERVER.SECRET);
        var err: string = await this.usecase.register(input, hashedPassword);

        if (err !== "") {
            var responseError = new ResponseError(400, err, {});
            responseError.send(res);
            return;
        }

        var responseSuccess = new ResponseSuccess(200, "User registered", {});
        responseSuccess.send(res);
        return;
    }

    private async login(req: any, res: any): Promise<void> {
        const input: ILoginRequest = req.body;

        var hashedPassword: string = hashPassword(input.password, this.config.SERVER.SECRET);
        var err: number = await this.usecase.checkUser(input, hashedPassword);
        
        if (err == -1) {
            var responseError = new ResponseError(400, "Can not find user.", {});
            responseError.send(res);
            return;
        }

        var user: IUser = {
            id: err,
            username: input.username,
            email: input.email,
            passwordHash: hashedPassword,
        }

        const accessToken = generateAccessToken(user, this.config);
        const refreshToken = generateRefreshToken(user, this.config);


        var responseSuccess = new ResponseSuccess(200, "User logged in", { accessToken, refreshToken });
        responseSuccess.send(res);
    }

    private async refresh(req: any, res: any): Promise<void> {
        var input: IRefreshTokenRequest = req.body;

        const decoded = verifyToken(input.refreshToken, this.config);

        if (decoded.id == undefined) {
            var responseError = new ResponseError(400, "Invalid token", {});
            responseError.send(res);
        }

        var user: IUser = {
            id: decoded.id,
            username: decoded.name,
            email: decoded.email,
            passwordHash: "",
        }

        const accessToken = generateAccessToken(user, this.config);
        const refreshToken = generateRefreshToken(user, this.config);

        var responseSuccess = new ResponseSuccess(200, "Token refreshed", { accessToken, refreshToken });
        responseSuccess.send(res);

    }

    public routes(router: any): void {
        // login, register, refresh
        router.post(this.routeString + "/register", async (req: any, res: any) => {
            await this.register(req, res);
        });

        router.post(this.routeString + "/login", async (req: any, res: any) => {
            await this.login(req, res);
        });

        router.post(this.routeString + "/refresh", async (req: any, res: any) => {
            await this.refresh(req, res);
        });
    }
}

export { ControllerAuth };

import { PrismaClient } from "@prisma/client";
import { IConfig } from "../core/commonServer/config";
import { IUsecasePing, newUsecasePing } from "../usecases/usecasePing";
import { IRepositoryPing, newRepositoryPing } from "../repositories/repositoryPing";
import express from 'express';
import { ResponseSuccess } from "../core/commonApi/responseSuccess";

class ControllerPing {
    public config: IConfig;
    public usecase: IUsecasePing;
    public routeString: string;

    constructor(config: IConfig, databaseClient: PrismaClient, routeString: string) {
        var repo: IRepositoryPing = newRepositoryPing(databaseClient);
        var usecase: IUsecasePing = newUsecasePing(repo);

        this.config = config;
        this.usecase = usecase;
        this.routeString = routeString;
    }

    public routes(router: express.Router): void{
        router.get(this.routeString + "/ping", (req, res) => {
            var response: ResponseSuccess = new ResponseSuccess(200,"pong",{});
            response.send(res);
        });

        // Create protected ping
    }
}

export { ControllerPing };
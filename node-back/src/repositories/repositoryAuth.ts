import { PrismaClient } from "@prisma/client";
import { ILoginRequest, IRegisterRequest } from "../core/commonApi/authApi";


interface IRepositoryAuth {
    dbClient: PrismaClient,
    register(input: IRegisterRequest, hashedPassword: string): Promise<string>;
    checkUser(input: ILoginRequest, hashedPassword: string): Promise<number>;
}

function newRepositoryAuth(dbClient: PrismaClient): IRepositoryAuth {
    return {
        dbClient,
        async register(input: IRegisterRequest, hashedPassword: string): Promise<string> {
            const { email, username } = input;
            try {
                const user = await this.dbClient.users.create({
                    data: {
                        email,
                        name: username,
                        password_hash: hashedPassword,
                    },
                });
                return "";
            } catch (error: any) {
                console.error("Error registering user:", error);
                return "Error while creating a user.";
            }
        },
        async checkUser(input: ILoginRequest, hashedPassword: string): Promise<number> {
            const {username, email } = input;
            try {
                const user = await this.dbClient.users.findUnique({
                    where: {
                        email: email,
                        password_hash: hashedPassword,
                    },
                });
                if (user) {
                    return user.id;
                }
                else {
                    throw new Error("User not found");
                }
            } catch (error: any) {
                console.error("Error checking user:", error);
                return -1;
            }
        },
        
            
    }

}

export {newRepositoryAuth, IRepositoryAuth};
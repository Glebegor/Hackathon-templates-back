import { PrismaClient } from "@prisma/client";
import { IConfig } from "../core/commonServer/config";
import express, { Request, Response } from 'express';
import { ControllerPing } from "../controllers/controllerPing";
import { ControllerAuth } from "../controllers/controllerAuth";

class IRouterMain {
    public config: IConfig;
    public databaseClient: PrismaClient;
    public router: express.Router;
    
    constructor(config: IConfig, databaseClient: PrismaClient) {
        this.config = config;
        this.databaseClient = databaseClient;
        this.router = express.Router();
    }

    public run(): express.Router {
        
        // Controllers
        var controllerPing: ControllerPing = new ControllerPing(this.config, this.databaseClient, '/test');
        var controllerAuth: ControllerAuth = new ControllerAuth(this.config, this.databaseClient, '/auth');
        
        // Routes
        controllerPing.routes(this.router);
        controllerAuth.routes(this.router);

        return this.router;
    }

}

function newRouterMain(config: IConfig, databaseClient: PrismaClient): IRouterMain {
    return new IRouterMain(config, databaseClient);
}

export { newRouterMain, IRouterMain };
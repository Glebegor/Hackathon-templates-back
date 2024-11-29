import { PrismaClient } from "@prisma/client";
import { IConfig } from "../core/commonServer/config"
import { Router } from "express";

function newControllerMain(config: IConfig, dbClient: PrismaClient): Router {
    const routerMain = Router();

    routerMain

    return routerMain;
}

export { newControllerMain };
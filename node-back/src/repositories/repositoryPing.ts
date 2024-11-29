import { PrismaClient } from "@prisma/client";

interface IRepositoryPing {
    dbClient: PrismaClient
}

function newRepositoryPing(dbClient: PrismaClient): IRepositoryPing {
    return {
        dbClient
    }
}

export {newRepositoryPing, IRepositoryPing};
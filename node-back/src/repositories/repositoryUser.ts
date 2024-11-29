import { PrismaClient } from "@prisma/client";


interface IRepositoryUser {
    dbClient: PrismaClient
}

function newRepositoryUser(dbClient: PrismaClient): IRepositoryUser {
    return {
        dbClient
    }
}

export {newRepositoryUser, IRepositoryUser};
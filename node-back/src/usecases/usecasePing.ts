import { IRepositoryPing } from "../repositories/repositoryPing";

interface IUsecasePing {
    repo: IRepositoryPing
}

function newUsecasePing(repo: IRepositoryPing): IUsecasePing {
    return {
        repo
    }
}

export {newUsecasePing, IUsecasePing};

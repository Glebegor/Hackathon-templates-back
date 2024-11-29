import {IRepositoryUser} from "../repositories/repositoryUser";

interface IUsecaseUser {
    repo: IRepositoryUser
}

function newUsecaseUser(repo: IRepositoryUser): IUsecaseUser {
    return {
        repo
    }
}

export {newUsecaseUser, IUsecaseUser};
import { ILoginRequest, IRegisterRequest } from "../core/commonApi/authApi";
import { IRepositoryAuth } from "../repositories/repositoryAuth";


interface IUsecaseAuth {
    repo: IRepositoryAuth
    register(input: IRegisterRequest, hashedPassword: string): Promise<string>;
    checkUser(input: ILoginRequest, hashedPassword: string): Promise<number>;
}

function newUsecaseAuth(repo: IRepositoryAuth): IUsecaseAuth {
    return {
        repo,
        register(input: IRegisterRequest, hashedPassword: string): Promise<string> {
            return repo.register(input, hashedPassword);
        },
        checkUser(input: ILoginRequest, hashedPassword: string): Promise<number> {
            return repo.checkUser(input, hashedPassword);
        }
    }
}


export { newUsecaseAuth, IUsecaseAuth };
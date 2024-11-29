import { IConfig } from "../core/commonServer/config";
import { IUsecasePing } from "../usecases/usecasePing";

interface IControllerPing {
    config: IConfig
    usecase: IUsecasePing
}

function newControllerPing(config: IConfig, usecase: IUsecasePing): IControllerPing {
    return {
        config,
        usecase
    }
}

export { newControllerPing, IControllerPing };
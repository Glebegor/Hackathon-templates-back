import * as dotenv from 'dotenv';
import { resolve } from 'path';
import { IConfig } from '../core/commonServer/config';

function getEnv(envType: string) {
    const envPath = resolve(__dirname, `../configs/env/.${envType}.env`);
    dotenv.config({ path: envPath });
}

function getConfig(envType: string) {
    getEnv(envType);
    const config: IConfig = require(`../configs/json/config.${envType}.json`);
    config.DB.PASSWORD = process.env.DB_PASSWORD;
    config.SERVER.SECRET = process.env.SERVER_SECRET;
    return config;
}

export default getConfig;

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
    
    process.env.DATABASE_URL = `postgresql://${config.DB.USER}:${config.DB.PASSWORD}@${config.DB.HOST}:${config.DB.PORT}/${config.DB.DATABASE}?schema=schema?sslmode=${config.DB.SSLMODE}`;
    return config;
}

export default getConfig;

import path from "path";
import getConfig from "./config";
import fs from "fs";
import { IConfig } from "../core/commonServer/config";

function setupEnv(envType: string): void {
    const config: IConfig = getConfig(envType);

    const databaseUrl = `postgresql://${config.DB.USER}:${config.DB.PASSWORD}@${config.DB.HOST}:${config.DB.PORT}/${config.DB.DATABASE}?sslmode=${config.DB.SSLMODE}`;
    process.env.DATABASE_URL = databaseUrl;

    const envFilePath = path.resolve(__dirname, `../../.env`);
    // clear whole file
    fs.writeFileSync(envFilePath, '');

    // add DATABASE_URL to .env file
    fs.appendFileSync(envFilePath, `DATABASE_URL="${databaseUrl}"\n`);

    console.log(`DATABASE_URL written to .env file: ${databaseUrl}`);
}

const args = process.argv.slice(2);

setupEnv(args[0] || 'dev');

export {setupEnv};
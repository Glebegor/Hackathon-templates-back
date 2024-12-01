import { IUser } from "../core/common/user";
import { IConfig } from "../core/commonServer/config";

const jwt = require('jsonwebtoken');

function generateAccessToken(user: IUser, config: IConfig) {
  return jwt.sign({ id: user.id, name: user.username, email: user.email}, config.SERVER.SECRET, { expiresIn: '60m' });
}

function generateRefreshToken(user: IUser, config: IConfig) {
  return jwt.sign({ id: user.id }, config.SERVER.SECRET, { expiresIn: '3d' });
}

function verifyToken(token: string, type = 'access', config: IConfig) {
    return jwt.verify(token, config.SERVER.SECRET);
}

export { generateAccessToken, generateRefreshToken, verifyToken };
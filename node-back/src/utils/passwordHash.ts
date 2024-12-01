import * as crypto from 'crypto';


function hashPassword(password: string, secretKey: any): string {
    const hash = crypto.createHmac('sha256', secretKey);
    hash.update(password);
    return hash.digest('hex');
}


export { hashPassword };
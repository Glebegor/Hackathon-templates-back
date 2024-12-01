

import { ResponseError } from "../core/commonApi/responseError";



function verityToken(req: any, res: any, next: any) {
    const bearerHeader = req.headers['authorization'];
    if (typeof bearerHeader !== 'undefined') {
        const bearer = bearerHeader.split(' ');
        const token = bearer[1];
        req.token = token;
        next();
    } else {
        var responseError = new ResponseError(403, "Forbidden", {});
        responseError.send(res);
    }    
}

export { verityToken };
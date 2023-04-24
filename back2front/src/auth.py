import jwt
from fastapi import Depends, HTTPException, status
from fastapi.security import APIKeyHeader

api_token_header = APIKeyHeader(name="Auth-Token")


def auth_header_token(token: str = Depends(api_token_header)) -> tuple[int, str]:
    try:
        info = jwt.decode(token, "secretKey", algorithms=["HS256"])
    except Exception as ex:
        print(ex)
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED)

    return int(info.get("data").get("id")), token

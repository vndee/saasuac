import os
import uvicorn
from loguru import logger
from fastapi import FastAPI, Form, File, Request
from fastapi.middleware.cors import CORSMiddleware


class UserAccessControlServer(object):
    server: FastAPI = FastAPI(
                title="User Access Control Server",
                contact={
                    "name": "Duy V. Huynh",
                    "email": "vndee.huynh@gmail.com"
                }
            )

    server.add_middleware(
                CORSMiddleware,
                allow_origins=["*"],
                allow_credentials=["*"],
                allow_methods=["*"],
                allow_headers=["*"]
            )
    
    @staticmethod
    def execute(host: str = "0.0.0.0", port: int = 8000):
        uvicorn.run(app=UserAccessControlServer.server, host=host, port=port, debug=True)

    @staticmethod
    @server.on_event("startup")
    async def startup_event():
        """doc
        """
        logger.info("UserAccessControlServer is living!!")

    @staticmethod
    @server.get("/healthcheck/")
    async def heathcheck():
        logger.info("Heathcheck request: Alive")
        return { "status": "alive", "message": "Hello, I'm here!" }


if __name__ == "__main__":
    server = UserAccessControlServer()
    server.execute()

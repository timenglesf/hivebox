"""
Main module for the FastAPI application.
"""

from fastapi import FastAPI

app = FastAPI()


@app.get("/")
async def root():
    """
    return {"message": "Hello World"}
    """
    return {"message": "Hello World"}

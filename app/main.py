"""
Main module for the FastAPI application.
"""

from fastapi import FastAPI
from requests import exceptions

from get_version import get_version_json
from get_temp import get_average_tempeature

app = FastAPI()


@app.get("/")
async def root():
    """
    return {"message": "Hello World"}
    """
    return {"message": "Hello World"}


@app.get("/version")
async def version():
    """
    Get the version of the application.
    """
    return get_version_json()


@app.get("/temperature")
async def temperature():
    """
    Get the average temperature of all sensebox sensors.
    """
    try:
        average_temperature = get_average_tempeature()
        return {"temperature": average_temperature, "unit": "°C"}
    except exceptions.RequestException as e:
        return {"error": e}

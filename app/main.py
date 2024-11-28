"""
Main module for the FastAPI application.
"""

from fastapi import FastAPI
from requests import exceptions
from pydantic import BaseModel

from get_version import get_version_json, Version
from get_temp import get_average_tempeature


app = FastAPI()


@app.get("/")
async def root():
    """
    return {"message": "Hello World"}
    """
    return {"message": "Hello World"}


@app.get("/version")
async def version() -> Version:
    """
    Get the version of the application.
    """
    return get_version_json()


class Temperature(BaseModel):
    """
    A class used to represent the temperature data.

    Attributes
    ----------
    temperature : float
        The temperature value.
    unit : str
        The unit of the temperature (e.g., "°C").
    """

    temperature: float
    unit: str


@app.get("/temperature")
async def temperature() -> Temperature | dict:
    """
    Get the average temperature of all sensebox sensors.
    """
    try:
        average_temperature = get_average_tempeature()
        return Temperature(temperature=average_temperature, unit="°C")
    except exceptions.RequestException as e:
        return {"error": str(e)}

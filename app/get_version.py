"""
Module for getting the current version of the application.
"""

from os import getenv
from main import app


def get_env_version():
    """
    Get the version of the application from the environment.
    """
    return getenv("VERSION")


def get_version_json():
    """
    Get the version of the application as a JSON object.
    """
    return {"version": get_env_version()}


@app.get("/version")
async def version():
    """
    Get the version of the application.
    """
    return get_version_json()

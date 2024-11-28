"""
Module to test the get_version.py module.
"""

from os import getenv
from fastapi.testclient import TestClient
from get_version import get_env_version, get_version_json, Version
from main import app

client = TestClient(app)


def test_get_env_version():
    """
    Test the get_env_version function
    """
    version = get_env_version()
    assert version is not None
    assert version == getenv("VERSION")


def test_get_version_json():
    """
    Test the get_version_json function
    """
    curr_version = get_env_version()
    json = get_version_json()
    assert json == Version(version=curr_version)


def test_version():
    """
    Test the /version endpoint.
    """
    response = client.get("/version")
    version = getenv("VERSION")
    assert response.status_code == 200
    assert response.json() == {"version": version}

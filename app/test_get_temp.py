"""Tests for the get_temp module."""

from datetime import datetime, timezone
from unittest.mock import patch, MagicMock
import pytest
from get_temp import (
    get_sensor_data,
    get_average_tempeature,
    format_time,
    filter_sensor_data,
    calculate_average_temperature,
)


@pytest.fixture
def mock_response():
    """Fixture for creating a mock response object."""
    return MagicMock()


@patch("get_temp.requests.get")
def test_get_sensor_data_success(mock_get, mock_res):
    """Test the successful retrieval of sensor data."""
    mock_res.status_code = 200
    mock_res.json.return_value = [
        {"sensors": [{"title": "temperature", "lastMeasurement": {"value": "20.5"}}]}
    ]
    mock_get.return_value = mock_response

    result = get_sensor_data()
    assert result == [
        {"sensors": [{"title": "temperature", "lastMeasurement": {"value": "20.5"}}]}
    ]


@patch("get_temp.requests.get")
def test_get_sensor_data_failure(mock_get, mock_res):
    """Test the failure scenario of retrieving sensor data."""
    mock_res.status_code = 400
    mock_res.json.return_value = {"message": "Bad Request"}
    mock_get.return_value = mock_response

    with pytest.raises(Exception):
        get_sensor_data()


def test_format_time():
    """Test the formatting of datetime objects to ISO 8601 strings."""

    time = datetime(2023, 10, 1, 12, 0, 0, tzinfo=timezone.utc)
    formatted_time = format_time(time)
    assert formatted_time == "2023-10-01T12:00:00Z"


def test_filter_sensor_data():
    """Test the filtering of sensor data to extract temperature values."""
    sensor_data = [
        {"sensors": [{"title": "temperature", "lastMeasurement": {"value": "20.5"}}]},
        {"sensors": [{"title": "Temperatur", "lastMeasurement": {"value": "21.5"}}]},
        {"sensors": [{"title": "humidity", "lastMeasurement": {"value": "50"}}]},
    ]
    result = filter_sensor_data(sensor_data)
    assert result == [20.5, 21.5]


def test_calculate_average_temperature():
    """Test the calculation of the average temperature."""
    temperatures = [20.5, 21.5, 19.5]
    result = calculate_average_temperature(temperatures)
    assert result == 20.5


@patch("get_temp.get_sensor_data")
@patch("get_temp.filter_sensor_data")
@patch("get_temp.calculate_average_temperature")
def test_get_average_tempeature(mock_calculate, mock_filter, mock_get):
    """Test the retrieval and calculation of the average temperature."""
    mock_get.return_value = [
        {"sensors": [{"title": "temperature", "lastMeasurement": {"value": "20.5"}}]},
        {"sensors": [{"title": "Temperatur", "lastMeasurement": {"value": "21.5"}}]},
    ]
    mock_filter.return_value = [20.5, 21.5]
    mock_calculate.return_value = 21.0

    result = get_average_tempeature()
    assert result == 21.0

    mock_get.side_effect = Exception("API Error")
    with pytest.raises(Exception):
        get_average_tempeature()

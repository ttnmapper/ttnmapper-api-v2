import requests
from celery import Celery

from ttnm_backend.settings import *

task_worker = Celery('api', backend='amqp', broker='amqp://guest@localhost//')

ERROR_WRONG_PARAMETERS = 15
ERROR_CONTACTING_OATH_PROVIDER = 16

@task_worker.task
def login_task(request: dict):
    """
    Exchange a code for a JWT token, update user details

    Will return a JSON object containing a result: "success/failure" If Failure, "error_code" and
    "error_message" will also exist.
    :param request: A dictionary of the request parameters. Should contain the code and state
    :return:
    """

    # Get the login code
    if "code" not in request:
        return failure(ERROR_WRONG_PARAMETERS, "Code parameter not supplied")

    code = request["code"]

    # Per specifications, there are no constraints placed on the code
    # https://tools.ietf.org/html/rfc6749#section-4.1.2
    # No documentation could be found on TTN side for the code.

    # Exchange code for Token
    data = {"grant-type": "authorization_code",
            "code": code,
            "redirect-uri": OAUTH_REDIRECT_URI,
            "client_id": OAUTH_CLIENT_ID}

    result = None
    try:
        result = requests.post(OAUTH_CODE_EXCHANGE_URL, data=data)
    except requests.exceptions.HTTPError as http_err:
        print(http_err)
        return failure(ERROR_CONTACTING_OATH_PROVIDER, "Error contacting the OATH provider")
    except requests.exceptions.ConnectionError as conn_err:
        print(conn_err)
        return failure(ERROR_CONTACTING_OATH_PROVIDER, "Error contacting the OATH provider")
    except requests.exceptions.Timeout as time_err:
        print(time_err)
        return failure(ERROR_CONTACTING_OATH_PROVIDER, "Error contacting the OATH provider")
    except requests.exceptions.RequestException as err:
        print(err)
        return failure(ERROR_CONTACTING_OATH_PROVIDER, "Error contacting the OATH provider")

    # How did this happen?
    if result is None:
        return {}

    # We have a token, and a refresh token
    token = result["token"]
    refresh_token = result["refresh_token"]

    # Get the user info

    return {"result": "success"}


def failure(error_code, error_message):
    """
    Helper function to create uniform error objects
    :param error_code:
    :param error_message:
    :return:
    """
    return {"result": "failure", "error_code": error_code, "error_message":error_message}
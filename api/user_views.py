from django.http import JsonResponse
"""
These views are used by user to log in and out
"""


def login_with_oauth_code(request):
    """"
    This is called after the user has logged in on TTN and accepted mapper login.
    The user is then returned to our 'Logging in' holding page while we update
    data on the user.

    The 'Logging In' page periodically polls the task until we have updated
    all data on the user.

    :return: String containing JWT token, used for login purposes
    """

    # Get the value of the redirect returned from the TTN site
    # token_type = request.GET.get('token_type', '')
    # access_token = request.GET.get('access_token', '')
    # expires_in = request.GET.get('expires_in')

    # We can now check the access_token for its info
    return JsonResponse({})


def check_login(request):
    """
    The 'Logging In' page checks the status of the task using this function.
    This function will return false until the task is successfull, upon which
    it will return the full jwt token
    :param request:
    :return:
    """

    return JsonResponse({})


def refresh_token(request):
    """
    Use a previously issued refresh token to get a new jwt
    :param request:
    :return:
    """
    print("refresh_token!")
    return JsonResponse({})

from django.test import TestCase
from api.tasks import login_task

import requests_mock
import ttnm_backend.settings


class LoginTaskTestCase(TestCase):
    def setUp(self):
        pass

    def test_framework_works(self):
        self.assertTrue(1)

    def test_wrong_request(self):
        """
        If a request contains the wrong keys it should be rejected
        :return:
        """
        # Check missing "code"
        request = {"invalid": "parameters"}
        response = login_task(request)

        self.assertEqual(response["result"], "failure")
        self.assertEqual(response["error_code"], 15)

    def test_perform_request(self):
        """

        :return:
        """

        request = {"code": "SECRET_CODE"}
        with requests_mock.mock() as m:
            m.post(ttnm_backend.settings.OAUTH_CODE_EXCHANGE_URL, text='data')
            response = login_task(request)
            print(response)

        self.assertTrue(False)
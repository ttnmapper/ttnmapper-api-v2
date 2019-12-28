from django.urls import path

from . import user_views

urlpatterns = [
    # path('', views.index, name='index'),
    path('login', user_views.login_with_oauth_code, name='login'),
    path('check_login', user_views.check_login, name='login'),
    path('refresh_token', user_views.refresh_token, name='login'),
]
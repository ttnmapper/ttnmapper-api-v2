from django.db import models

from django.utils import timezone

# Create your models here.

USER_STATE_INACTIVE = 0     # Account has been deactivated
USER_STATE_ACTIVE = 1       # User is allowed to login


class User(models.Model):

    # The user GUID is used internally to reference a user.
    guid = models.guid()

    # The username is the string used to reference the user (i.e in the login box / user page)
    username = models.Charfield(max_length=200)

    # firstSeen stores the date we first saw the user.
    firstSeen = models.DateTimeField(default=timezone.now)

    # Stores the last time the user logged in
    lastSeen = models.DateTimeField(default=timezone.now)

    # activeState indicates if the user is active/inactive. Stored as int if we need to add mroe later
    activeState = USER_STATE_ACTIVE



package opa.authz

import input.user_id as user_id
import input.request as request
import input.user_roles as user_roles

default allow = false


allow{
    # request.method = "POST"
    request.path = ["proto.AccountService","CreateBankEmployee"]
    allowed_roles=["admin"]
    allowed_roles[_] = user_roles[_]
    user_id
}

allow{
    # request.method = "POST"
    request.path = ["proto.AccountService","GetEmployee"]
    allowed_roles=["admin"]
    allowed_roles[_] = user_roles[_]
    user_id
}
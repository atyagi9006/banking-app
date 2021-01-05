package opa.authz

test_create_employee{
    allow with input as {"request": {"path": ["proto.AccountService","CreateBankEmployee"]},"user_roles": {"abc@xyz.com-":"admin"}, "user_id": "abc@xyz.com-"}
    not allow with input as {"request": {"path": ["proto.AccountService","CreateBankEmployee"]},"user_roles": {"abc@xyz.com-":"staff"}, "user_id": "abc@xyz.com-"}
    not allow with input as {"request": {"path": ["proto.AccountService","CreateBankEmployee"]},"user_roles": {"abc@xyz.com-":"admin"}}
    not allow with input as {"request": {"path": ["proto.AccountService","CreateBankEmployee"]}}
}

test_get_employee{
    allow with input as {"request": {"path": ["proto.AccountService","GetEmployee"]},"user_roles": {"abc@xyz.com-":"admin"}, "user_id": "abc@xyz.com-"}
    not allow with input as {"request": {"path": ["proto.AccountService","GetEmployee"]},"user_roles": {"abc@xyz.com-":"staff"}, "user_id": "abc@xyz.com-"}
    not allow with input as {"request": {"path": ["proto.AccountService","GetEmployee"]},"user_roles": {"abc@xyz.com-":"admin"}}
    not allow with input as {"request": {"path": ["proto.AccountService","GetEmployee"]}}
}
package user

var insUser = `
INSERT INTO user_info (phone_number, password) VALUES (?, ?);
`

var selUser = `
SELECT id, password FROM user_info WHERE phone_number = ?;
`

var insAccess = `
INSERT INTO access_control (user_id, access_token, expires_at)
VALUES (?, ?, DATE_ADD(NOW(), INTERVAL 10 MINUTE))
ON DUPLICATE KEY UPDATE
access_token = VALUES(access_token),
expires_at = DATE_ADD(NOW(), INTERVAL 10 MINUTE),
created_at = NOW();
`

var delAccess = `
DELETE FROM access_control WHERE user_id = ?;
`

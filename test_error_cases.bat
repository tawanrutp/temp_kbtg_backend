@echo off
REM Advanced test script with error cases
REM Make sure the server is running before executing this script

echo ============================================
echo Transfer API - Error Handling Test Script
echo ============================================
echo.

echo [Test 1] Attempting transfer with insufficient balance...
echo Creating user with only 50 points...
curl -X POST http://localhost:3000/api/v1/users -H "Content-Type: application/json" -d "{\"name\":\"Charlie\",\"email\":\"charlie@example.com\",\"balance\":50}"
echo.
echo Attempting to transfer 100 points (should fail)...
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":3,\"to_user_id\":1,\"amount\":100,\"note\":\"Should fail\",\"idempotency_key\":\"transfer-error-001\"}"
echo.
echo.

timeout /t 2 >nul

echo [Test 2] Attempting self-transfer (should fail)...
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":1,\"amount\":10,\"note\":\"Self transfer\",\"idempotency_key\":\"transfer-error-002\"}"
echo.
echo.

timeout /t 2 >nul

echo [Test 3] Attempting transfer with non-existent user...
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":999,\"amount\":10,\"note\":\"Invalid user\",\"idempotency_key\":\"transfer-error-003\"}"
echo.
echo.

timeout /t 2 >nul

echo [Test 4] Attempting transfer with invalid amount...
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":-10,\"note\":\"Negative amount\",\"idempotency_key\":\"transfer-error-004\"}"
echo.
echo.

timeout /t 2 >nul

echo [Test 5] Attempting transfer without idempotency_key...
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":10,\"note\":\"No idempotency key\"}"
echo.
echo.

timeout /t 2 >nul

echo [Test 6] Getting non-existent transfer...
curl http://localhost:3000/api/v1/transfers/non-existent-key
echo.
echo.

timeout /t 2 >nul

echo [Test 7] Create a successful transfer and then try to cancel...
echo Creating transfer...
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":10,\"note\":\"Test cancel\",\"idempotency_key\":\"transfer-cancel-001\"}"
echo.
echo Attempting to cancel completed transfer (should fail)...
curl -X DELETE http://localhost:3000/api/v1/transfers/transfer-cancel-001
echo.
echo.

echo ============================================
echo Error handling tests completed!
echo ============================================
echo.
echo Summary of expected behaviors:
echo [Test 1] 400 - Insufficient balance
echo [Test 2] 400 - Cannot transfer to same user
echo [Test 3] 404 - User not found
echo [Test 4] 400 - Invalid amount
echo [Test 5] 400 - Missing required fields
echo [Test 6] 404 - Transfer not found
echo [Test 7] 400 - Cannot cancel completed transfer
echo.

pause

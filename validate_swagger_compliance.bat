@echo off
REM Swagger Validation Script
REM This script tests if the API implementation matches the swagger.yml specification

echo ============================================
echo Swagger Specification Validation
echo ============================================
echo.
echo This script will test the API against the swagger.yml specification
echo Make sure the server is running on http://localhost:3000
echo.
pause

echo [Step 1] Creating test users...
echo.

REM Create User 1
echo Creating User 1 (Alice)...
curl -X POST http://localhost:3000/api/v1/users -H "Content-Type: application/json" -d "{\"name\":\"Alice\",\"email\":\"alice-test@example.com\",\"balance\":1000}" > response1.json
type response1.json
echo.
echo Expected: 201 Created with success=true and data containing user object
echo.
timeout /t 2 >nul

REM Create User 2
echo Creating User 2 (Bob)...
curl -X POST http://localhost:3000/api/v1/users -H "Content-Type: application/json" -d "{\"name\":\"Bob\",\"email\":\"bob-test@example.com\",\"balance\":500}" > response2.json
type response2.json
echo.
echo Expected: 201 Created with success=true and data containing user object
echo.
timeout /t 2 >nul

echo.
echo [Step 2] Testing GET endpoints...
echo.

REM Get all users
echo GET /api/v1/users
curl http://localhost:3000/api/v1/users > response3.json
type response3.json
echo.
echo Expected: 200 OK with success=true and array of users
echo.
timeout /t 2 >nul

REM Get user by ID
echo GET /api/v1/users/1
curl http://localhost:3000/api/v1/users/1 > response4.json
type response4.json
echo.
echo Expected: 200 OK with success=true and user object
echo.
timeout /t 2 >nul

REM Get user balance
echo GET /api/v1/users/1/balance
curl http://localhost:3000/api/v1/users/1/balance > response5.json
type response5.json
echo.
echo Expected: 200 OK with success=true and data containing user_id and balance
echo.
timeout /t 2 >nul

echo.
echo [Step 3] Testing transfer creation...
echo.

REM Create transfer
echo POST /api/v1/transfers
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":100,\"note\":\"Test transfer\",\"idempotency_key\":\"swagger-validation-001\"}" > response6.json
type response6.json
echo.
echo Expected: 201 Created with success=true and transfer object with status=completed
echo.
timeout /t 2 >nul

echo.
echo [Step 4] Testing idempotency...
echo.

REM Test idempotency (same request)
echo POST /api/v1/transfers (same idempotency_key)
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":100,\"note\":\"Test transfer\",\"idempotency_key\":\"swagger-validation-001\"}" > response7.json
type response7.json
echo.
echo Expected: 200 OK with same transfer and message "Transfer already exists (idempotent)"
echo.
timeout /t 2 >nul

echo.
echo [Step 5] Testing transfer queries...
echo.

REM Get all transfers
echo GET /api/v1/transfers
curl http://localhost:3000/api/v1/transfers > response8.json
type response8.json
echo.
echo Expected: 200 OK with success=true and array of transfers
echo.
timeout /t 2 >nul

REM Get transfer by idempotency_key
echo GET /api/v1/transfers/swagger-validation-001
curl http://localhost:3000/api/v1/transfers/swagger-validation-001 > response9.json
type response9.json
echo.
echo Expected: 200 OK with success=true and transfer object
echo.
timeout /t 2 >nul

REM Filter transfers by user_id
echo GET /api/v1/transfers?user_id=1
curl http://localhost:3000/api/v1/transfers?user_id=1 > response10.json
type response10.json
echo.
echo Expected: 200 OK with filtered transfers
echo.
timeout /t 2 >nul

REM Filter transfers by status
echo GET /api/v1/transfers?status=completed
curl http://localhost:3000/api/v1/transfers?status=completed > response11.json
type response11.json
echo.
echo Expected: 200 OK with filtered transfers
echo.
timeout /t 2 >nul

echo.
echo [Step 6] Testing ledger endpoint...
echo.

REM Get user ledger
echo GET /api/v1/users/1/ledger
curl http://localhost:3000/api/v1/users/1/ledger > response12.json
type response12.json
echo.
echo Expected: 200 OK with success=true and array of ledger entries
echo.
timeout /t 2 >nul

REM Filter ledger by event_type
echo GET /api/v1/users/1/ledger?event_type=transfer_out
curl http://localhost:3000/api/v1/users/1/ledger?event_type=transfer_out > response13.json
type response13.json
echo.
echo Expected: 200 OK with filtered ledger entries
echo.
timeout /t 2 >nul

echo.
echo [Step 7] Testing error cases...
echo.

REM Test 404 - User not found
echo GET /api/v1/users/999 (should return 404)
curl -i http://localhost:3000/api/v1/users/999 2>&1 | find "404"
echo Expected: 404 Not Found with error message
echo.
timeout /t 2 >nul

REM Test 400 - Invalid transfer (same user)
echo POST /api/v1/transfers with same from_user_id and to_user_id (should return 400)
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":1,\"amount\":10,\"idempotency_key\":\"swagger-error-001\"}" > response14.json
type response14.json
echo.
echo Expected: 400 Bad Request with error "Cannot transfer to the same user"
echo.
timeout /t 2 >nul

REM Test 400 - Insufficient balance
echo POST /api/v1/transfers with insufficient balance (should return 400)
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":2,\"to_user_id\":1,\"amount\":10000,\"idempotency_key\":\"swagger-error-002\"}" > response15.json
type response15.json
echo.
echo Expected: 400 Bad Request with error "Insufficient balance"
echo.
timeout /t 2 >nul

REM Test 400 - Invalid amount (negative)
echo POST /api/v1/transfers with negative amount (should return 400)
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":-10,\"idempotency_key\":\"swagger-error-003\"}" > response16.json
type response16.json
echo.
echo Expected: 400 Bad Request with error message
echo.
timeout /t 2 >nul

REM Test 400 - Missing required field
echo POST /api/v1/transfers without idempotency_key (should return 400)
curl -X POST http://localhost:3000/api/v1/transfers -H "Content-Type: application/json" -d "{\"from_user_id\":1,\"to_user_id\":2,\"amount\":10}" > response17.json
type response17.json
echo.
echo Expected: 400 Bad Request with error "Missing required fields or invalid amount"
echo.
timeout /t 2 >nul

REM Test 404 - Transfer not found
echo GET /api/v1/transfers/non-existent-key (should return 404)
curl http://localhost:3000/api/v1/transfers/non-existent-key > response18.json
type response18.json
echo.
echo Expected: 404 Not Found with error "Transfer not found"
echo.
timeout /t 2 >nul

echo.
echo [Step 8] Testing UPDATE and DELETE...
echo.

REM Update user
echo PUT /api/v1/users/1
curl -X PUT http://localhost:3000/api/v1/users/1 -H "Content-Type: application/json" -d "{\"name\":\"Alice Updated\",\"email\":\"alice-test@example.com\",\"balance\":900}" > response19.json
type response19.json
echo.
echo Expected: 200 OK with updated user object
echo.
timeout /t 2 >nul

REM Try to cancel completed transfer (should fail)
echo DELETE /api/v1/transfers/swagger-validation-001 (should fail - already completed)
curl -X DELETE http://localhost:3000/api/v1/transfers/swagger-validation-001 > response20.json
type response20.json
echo.
echo Expected: 400 Bad Request with error "Cannot cancel transfer with status: completed"
echo.
timeout /t 2 >nul

echo.
echo ============================================
echo Validation Complete!
echo ============================================
echo.
echo Please review the responses above to verify they match the swagger.yml specification:
echo.
echo ✓ Response format: {"success": true, "data": {...}} for success
echo ✓ Response format: {"error": "..."} for errors
echo ✓ Status codes: 200 (GET), 201 (POST create), 400 (bad request), 404 (not found)
echo ✓ Required fields present in responses
echo ✓ Data types match specification
echo ✓ Idempotency working correctly
echo ✓ Validation rules enforced
echo ✓ Error messages clear and descriptive
echo.
echo All response files saved as response1.json to response20.json for detailed review
echo.

REM Cleanup
echo Cleaning up response files...
del response*.json

pause

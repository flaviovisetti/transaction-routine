#!/bin/bash

psql -h db -d transaction_routine -U postgres -c "INSERT INTO operation_types(description) VALUES('COMPRA A VISTA');"
psql -h db -d transaction_routine -U postgres -c "INSERT INTO operation_types(description) VALUES('COMPRA PARCELADA');"
psql -h db -d transaction_routine -U postgres -c "INSERT INTO operation_types(description) VALUES('SAQUE');"
psql -h db -d transaction_routine -U postgres -c "INSERT INTO operation_types(description) VALUES('PAGAMENTO');"
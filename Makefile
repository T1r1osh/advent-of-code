# Variables
MAIN_DIR := solutions
GO_CMD := go run
TEST_CMD := go test


# Target to test the program
test-%:
	@YEAR=$(shell echo $@ | cut -c6-7) \
	&& DAY=$(shell echo $@ | cut -c8-9) \
	&& FULL_YEAR="20$$YEAR" \
	&& FULL_DAY="Day$$DAY" \
	&& cd $(MAIN_DIR)/$$FULL_YEAR/"$$FULL_DAY"/ \
	&& $(TEST_CMD) 

# Target for dynamic execution based on year and day
%:
	@YEAR=$(shell echo $@ | cut -c1-2) \
	&& DAY=$(shell echo $@ | cut -c3-4) \
	&& FULL_YEAR="20$$YEAR" \
	&& FULL_DAY="Day$$DAY" \
	&& $(GO_CMD) $(MAIN_DIR)/$$FULL_YEAR/"$$FULL_DAY"/*.go


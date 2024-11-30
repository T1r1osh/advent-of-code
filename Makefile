# Variables
MAIN_DIR := solutions
GO_CMD := go run

# Target for dynamic execution based on year and day
%:
	@YEAR=$(shell echo $@ | cut -c1-2) \
	&& DAY=$(shell echo $@ | cut -c3-4) \
	&& FULL_YEAR="20$$YEAR" \
	&& FULL_DAY="Day $$DAY" \
	&& $(GO_CMD) $(MAIN_DIR)/$$FULL_YEAR/"$$FULL_DAY"/main.go

package format

import (
	"fmt"
	"strconv"
)

const ID_SIZE int = 2
const LEN_SIZE int = 2

func parsedData(data string, index *int) (string, string, string, error) {
	if len(data) < ID_SIZE+LEN_SIZE {
		return "", "", "", fmt.Errorf("data too short")
	}
	id := data[*index : *index+ID_SIZE]
	*index += ID_SIZE
	payloadLenStr := data[*index : *index+LEN_SIZE]
	*index += LEN_SIZE
	payloadLen, err := strconv.Atoi(payloadLenStr)
	if err != nil {
		return "", "", "", fmt.Errorf("invalid payload length %s", payloadLenStr)
	}
	if len(data) < ID_SIZE+LEN_SIZE+payloadLen {
		return "", "", "", fmt.Errorf("data too short for payload")
	}
	payload := data[*index : *index+payloadLen]
	*index += payloadLen
	return id, payloadLenStr, payload, nil
}

func formatCaseOutOfList(id int, payload string) (string, error) {
	var response string
	for i := 0; i < len(payload); {
		subID, payloadLenStr, payloadData, err := parsedData(payload, &i)
		if err != nil {
			return "", fmt.Errorf("error parsing payload: %v", err)
		}
		response += fmt.Sprintf(". . . %s %s", subID, payloadLenStr)
		if id == 38 && subID == "01" {
			response += "\n"
			for j := i - len(payloadData); j < i; {
				subSubID, subPayloadLenStr, subPayloadData, err := parsedData(payload, &j)
				if err != nil {
					return "", fmt.Errorf("error parsing sub-payload: %v", err)
				}
				response += fmt.Sprintf(". . . . . . %s %s %s\n", subSubID, subPayloadLenStr, subPayloadData)
			}
		} else {
			response += fmt.Sprintf(" %s\n", payloadData)
		}
	}
	return response, nil
}

func FormatQR(data string) (string, error) {
	// Placeholder for QR code formatting logic
	// In a real implementation, this function would format the QR code data
	// according to specific requirements or standards.
	var formattedData string
	if data == "" {
		return "", fmt.Errorf("data cannot be empty")
	}
	for i := 0; i < len(data); {
		idStr, payloadLenStr, payload, err := parsedData(data, &i)
		if err != nil {
			return "", fmt.Errorf("error parsing data at index %d: %v", i, err)
		}
		id, _ := strconv.Atoi(idStr)
		switch {
		case id >= 26 && id <= 51:
			payloadFormatted, err := formatCaseOutOfList(id, payload)
			if err != nil {
				return "", fmt.Errorf("error formatting payload of id %s: %v", idStr, err)
			}
			formattedData += fmt.Sprintf("%s %s\n%s", idStr, payloadLenStr, payloadFormatted)
		default:
			formattedData += fmt.Sprintf("%s %s %s\n", idStr, payloadLenStr, payload)
		}
	}
	return formattedData, nil
}

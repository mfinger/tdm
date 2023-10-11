package TasDevMgr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

type Device struct {
	Name    string
	Address net.IP
	Version string
	Config  []byte
}

func (device Device) send(path string) ([]byte, error) {
	url := fmt.Sprintf("http://%s/%s", device.Address.String(), path)
	var err error
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (device Device) sendCommand(command string) (map[string]interface{}, error) {
	if data, err := device.send(fmt.Sprintf("cm?cmnd=%s", command)); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to send command to device: %s", err.Error()))
	} else {
		var response map[string]interface{}
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, errors.New(fmt.Sprintf("Failed to unmarshal command response responsee: %s", err.Error()))
		}
		return response, nil
	}
}

func (device *Device) Probe() error {
	if response, err := device.sendCommand("Status"); err != nil {
		return errors.New(fmt.Sprintf("Failed to get device status: %s", err.Error()))
	} else {
		var status map[string]interface{}
		if val, found := response["Status"]; found {
			status = val.(map[string]interface{})
		} else {
			return errors.New("device return unexpected response, no 'Status' element found")
		}
		device.Name = status["DeviceName"].(string)

		if response, err := device.sendCommand("Status+2"); err != nil {
			return errors.New(fmt.Sprintf("Failed to get device firmware status: %s", err.Error()))
		} else {
			var status map[string]interface{}
			if val, found := response["StatusFWR"]; found {
				status = val.(map[string]interface{})
			} else {
				return errors.New("device return unexpected response, no 'Status' element found")
			}
			device.Version = status["Version"].(string)

		}
		return nil
	}
}

func (device Device) Print() {
	fmt.Printf("%-20s %-10s %s\n", device.Name, device.Address.String(), device.Version)
}

func (device Device) SaveConfig(filename string) error {
	if err := os.WriteFile(filename, device.Config, 0644); err != nil {
		return err
	}
	return nil
}

func (device *Device) LoadConfig(filename string) error {
	if data, err := os.ReadFile(filename); err != nil {
		return err
	} else {
		device.Config = data
	}
	return nil
}

func (device *Device) FetchConfig() error {

	if data, err := device.send("dl?"); err != nil {
		return err
	} else {
		device.Config = data
		return nil
	}
}

func (device Device) SendConfig() {

}

func (device Device) Backup(filename string) error {
	if device.Address == nil {
		return errors.New("Invalid device address")
	}

	if err := device.FetchConfig(); err != nil {
		return err
	}
	if err := os.WriteFile(filename, device.Config, 0644); err != nil {
		return errors.New(fmt.Sprintf("Failed to fetch config: %s", err.Error()))
	}

	if err := device.SaveConfig(filename); err != nil {
		return errors.New(fmt.Sprintf("Failed to save config: %s", err.Error()))
	}

	return nil
}

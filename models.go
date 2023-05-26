package IMEKeyboard

import (
	sf "github.com/sa-/slicefunk"
)

// BrandNameYamlModel Get the input-method brand name of device. Different device size have different
// input-method size and start y-loc
// Field DeviceName: the name of device. e.g:
// Google Pixel 3, Google Pixel XL 2...
// Width Device Screen Width
type BrandNameYamlModel struct {
	DeviceName string `yaml:"name"`

	SogouStartY uint64 `yaml:"sogou"`

	BaiduStartY uint64 `yaml:"baidu"`

	IflytekStartY uint64 `yaml:"iflytek"`

	HWBaiduStartY uint64 `yaml:"hwbaidu"`

	WeChatStartY uint64 `yaml:"wechat"`

	Width uint64 `yaml:"width"`
}

// IsDevice Confirm the device is the same in yaml devices
func (model *BrandNameYamlModel) IsDevice(deviceName string) bool {
	return deviceName == model.DeviceName
}

// ScreenWidthYamlModel Get the specific elements of different width devices
type ScreenWidthYamlModel struct {
	LogoHeight uint64 `yaml:"logo_height"`

	MenuHeight uint64 `yaml:"menu_height"`

	KeyboardHeight uint64 `yaml:"keyboard_height"`

	SlideX uint64 `yaml:"slide_x"`

	SlideY uint64 `yaml:"slide_y"`

	DownArrowWidth uint64 `yaml:"down_arrow_width"`
}

// BrandDetailPosModel Get the specific elements of different width devices
type BrandDetailPosModel struct {
	Width uint64 `yaml:"width"`

	Sogou ScreenWidthYamlModel `yaml:"sogou"`

	Baidu ScreenWidthYamlModel `yaml:"baidu"`

	Iflytek ScreenWidthYamlModel `yaml:"iflytek"`

	HWBaidu ScreenWidthYamlModel `yaml:"hwbaidu"`

	WeChat ScreenWidthYamlModel `yaml:"wechat"`
}

// IsWidth Filter the same width in the screen
func (model BrandDetailPosModel) IsWidth(width uint64) bool {
	return width == model.Width
}

type MobileYamlReader struct {
	BrandNames []BrandNameYamlModel `yaml:"brand_name"`

	WidthConfig []BrandDetailPosModel `yaml:"width"`
}

type MobileYamlModel struct {
	Mobile MobileYamlReader `yaml:"mobile"`
}

// KeyboardRealConfigModel the keyboard element of real device
type KeyboardRealConfigModel struct {
	SogouStartY uint64

	BaiduStartY uint64

	IflytekStartY uint64

	HWBaiduStartY uint64

	WeChatStartY uint64

	BrandDetailModel BrandDetailPosModel
}

type KeyboardBrandType int64

const (
	SogouBrand KeyboardBrandType = iota
	BaiduBrand
	HWBaiduBrand
	IflytekBrand
	WeChatBrand
)

type KeyboardType int64

const (
	Keyboard26Key KeyboardType = iota
	Keyboard9Key
)

type KeyboardAxiosPos struct {
	X int64
	Y int64
}

// ImeStartY return the input-method keyboard start y pos
func (model *KeyboardRealConfigModel) ImeStartY(keyboardType KeyboardBrandType) uint64 {
	switch keyboardType {
	case SogouBrand:
		return model.SogouStartY
	case BaiduBrand:
		return model.BaiduStartY
	case HWBaiduBrand:
		return model.HWBaiduStartY
	case WeChatBrand:
		return model.WeChatStartY
	case IflytekBrand:
		return model.IflytekStartY
	}
	return 0
}

// GetMenuY get the keyboard menu height
func (model *KeyboardRealConfigModel) GetMenuY(keyboardType KeyboardBrandType) uint64 {
	var ret uint64 = 0
	switch keyboardType {
	case SogouBrand:
		ret = model.SogouStartY + model.BrandDetailModel.Sogou.LogoHeight + model.BrandDetailModel.Sogou.MenuHeight/2
	case BaiduBrand:
		ret = model.BaiduStartY + model.BrandDetailModel.Baidu.LogoHeight + model.BrandDetailModel.Baidu.MenuHeight/2
	case HWBaiduBrand:
		ret = model.HWBaiduStartY + model.BrandDetailModel.HWBaidu.MenuHeight/2
	case WeChatBrand:
		ret = model.WeChatStartY + model.BrandDetailModel.WeChat.MenuHeight/2 + model.BrandDetailModel.WeChat.LogoHeight
	case IflytekBrand:
		ret = model.IflytekStartY + model.BrandDetailModel.Iflytek.MenuHeight/2 + model.BrandDetailModel.Iflytek.LogoHeight
	}
	return ret
}

// GetCandidateHeight get the height of candidate view
func (model *KeyboardRealConfigModel) GetCandidateHeight(keyboardType KeyboardBrandType) uint64 {
	var ret uint64 = 0
	switch keyboardType {
	case SogouBrand:
		ret = model.BrandDetailModel.Sogou.LogoHeight + model.BrandDetailModel.Sogou.MenuHeight
	case BaiduBrand:
		ret = model.BrandDetailModel.Baidu.LogoHeight + model.BrandDetailModel.Baidu.MenuHeight
	case HWBaiduBrand:
		ret = model.BrandDetailModel.HWBaidu.MenuHeight
	case WeChatBrand:
		ret = model.BrandDetailModel.WeChat.LogoHeight + model.BrandDetailModel.WeChat.MenuHeight
	case IflytekBrand:
		ret = model.BrandDetailModel.Iflytek.LogoHeight + model.BrandDetailModel.Iflytek.MenuHeight
	}
	return ret
}

// GetSingleMenuWidth get the single menu key width
func (model *KeyboardRealConfigModel) GetSingleMenuWidth(keyboardType KeyboardBrandType) uint64 {
	var ret uint64 = 0
	switch keyboardType {
	case SogouBrand:
		ret = (model.BrandDetailModel.Width - model.BrandDetailModel.Sogou.DownArrowWidth) / 6
	case BaiduBrand:
		ret = (model.BrandDetailModel.Width - model.BrandDetailModel.Baidu.DownArrowWidth) / 6
	case HWBaiduBrand:
		ret = model.BrandDetailModel.HWBaidu.DownArrowWidth
	case WeChatBrand:
		ret = model.BrandDetailModel.WeChat.DownArrowWidth
	case IflytekBrand:
		ret = (model.BrandDetailModel.Width - model.BrandDetailModel.Iflytek.DownArrowWidth) / 7
	}
	return ret
}

func (model *KeyboardRealConfigModel) GetKeyElementWidth(keyboardType KeyboardBrandType) uint64 {
	var ret uint64 = 0
	switch keyboardType {
	case SogouBrand:
		ret = (model.BrandDetailModel.Width - 2*model.BrandDetailModel.Sogou.SlideX) / 10
	case BaiduBrand:
		ret = (model.BrandDetailModel.Width - 2*model.BrandDetailModel.Baidu.SlideX) / 10
	case HWBaiduBrand:
		ret = (model.BrandDetailModel.Width - 2*model.BrandDetailModel.HWBaidu.SlideX) / 10
	case WeChatBrand:
		ret = model.BrandDetailModel.Width / 10
	case IflytekBrand:
		ret = (model.BrandDetailModel.Width - 2*model.BrandDetailModel.Iflytek.SlideX) / 10
	}
	return ret
}

// GetKeyElementHeight get keyboard element height
func (model *KeyboardRealConfigModel) GetKeyElementHeight(keyboardType KeyboardBrandType) uint64 {
	var ret uint64 = 0
	switch keyboardType {
	case SogouBrand:
		ret = (model.BrandDetailModel.Sogou.KeyboardHeight - 2*model.BrandDetailModel.Sogou.SlideY) / 4
	case BaiduBrand:
		ret = (model.BrandDetailModel.Baidu.KeyboardHeight - 2*model.BrandDetailModel.Baidu.SlideY) / 4
	case HWBaiduBrand:
		ret = (model.BrandDetailModel.HWBaidu.KeyboardHeight - 2*model.BrandDetailModel.HWBaidu.SlideY) / 4
	case WeChatBrand:
		ret = (model.BrandDetailModel.WeChat.KeyboardHeight - 2*model.BrandDetailModel.WeChat.SlideY) / 4
	case IflytekBrand:
		ret = (model.BrandDetailModel.Iflytek.KeyboardHeight - 2*model.BrandDetailModel.Iflytek.SlideY) / 4
	}
	return ret
}

// GetKeyboardStartY get the y axios of keyboard key elements
func (model *KeyboardRealConfigModel) GetKeyboardStartY(keyboardType KeyboardBrandType) uint64 {
	var ret uint64 = 0
	switch keyboardType {
	case SogouBrand:
		ret = model.SogouStartY + model.BrandDetailModel.Sogou.LogoHeight + model.BrandDetailModel.Sogou.MenuHeight + model.BrandDetailModel.Sogou.SlideY
	case BaiduBrand:
		ret = model.BaiduStartY + model.BrandDetailModel.Baidu.LogoHeight + model.BrandDetailModel.Baidu.MenuHeight + model.BrandDetailModel.Baidu.SlideY
	case HWBaiduBrand:
		ret = model.HWBaiduStartY + model.BrandDetailModel.HWBaidu.MenuHeight + model.BrandDetailModel.HWBaidu.SlideY
	case WeChatBrand:
		ret = model.WeChatStartY + model.BrandDetailModel.WeChat.MenuHeight + model.BrandDetailModel.WeChat.LogoHeight
	case IflytekBrand:
		ret = model.IflytekStartY + model.BrandDetailModel.Iflytek.LogoHeight + model.BrandDetailModel.Iflytek.MenuHeight + model.BrandDetailModel.Iflytek.SlideY
	}
	return ret
}

// DetailMenuPos get the detail keyboard pos menu
func (model *KeyboardRealConfigModel) DetailMenuPos(keyboardType KeyboardBrandType) *KeyboardAxiosPos {
	var menuYAxios = int64(model.GetMenuY(keyboardType))
	var menuXAxios int64
	if keyboardType == HWBaiduBrand || keyboardType == WeChatBrand {
		menuXAxios = int64(model.GetSingleMenuWidth(keyboardType) / 2)
	} else {
		menuXAxios = int64(model.GetSingleMenuWidth(keyboardType))
	}
	return &KeyboardAxiosPos{
		X: menuXAxios,
		Y: menuYAxios,
	}
}

// Nine9KeyWidth get the keyboard width
func (model *KeyboardRealConfigModel) Nine9KeyWidth(keyboardType KeyboardBrandType) uint64 {
	startX := model.GetKeyElementWidth(keyboardType) * 16 / 10
	return (model.BrandDetailModel.Width - startX*2) / 3
}

// MenuPos the position of menu
func (model *KeyboardRealConfigModel) MenuPos(keyboardType KeyboardBrandType) *KeyboardAxiosPos {
	var menuYAxios int64
	var menuXAxios int64
	switch keyboardType {
	case IflytekBrand:
		menuXAxios = int64(model.GetSingleMenuWidth(keyboardType)*24/10 + 10)
		menuYAxios = int64(model.GetMenuY(keyboardType))
	case HWBaiduBrand:
		menuXAxios = int64(model.GetSingleMenuWidth(keyboardType)*2 + model.BrandDetailModel.HWBaidu.SlideX)
		menuYAxios = int64(model.GetKeyboardStartY(keyboardType) + model.GetKeyElementHeight(keyboardType)*15/10)
	case WeChatBrand:
		menuXAxios = int64(model.BrandDetailModel.Width / 8 * 5)
		menuYAxios = int64(model.GetMenuY(keyboardType)) + int64(model.BrandDetailModel.WeChat.MenuHeight/2) + int64(model.BrandDetailModel.Width/6)
	default:
		menuXAxios = int64(model.GetSingleMenuWidth(keyboardType) * 25 / 10)
		menuYAxios = int64(model.GetMenuY(keyboardType))
	}
	return &KeyboardAxiosPos{
		X: menuXAxios,
		Y: menuYAxios,
	}
}

// ChNineKeyEle Chinese keyboard element key
func (model *KeyboardRealConfigModel) ChNineKeyEle(keyboardType KeyboardBrandType) *KeyboardAxiosPos {
	var menuYAxios int64
	menuXAxios := int64(model.GetKeyElementWidth(keyboardType) * 15 / 10)
	if keyboardType == HWBaiduBrand || keyboardType == WeChatBrand {
		menuYAxios = int64(model.GetKeyElementHeight(keyboardType)*12/10 + model.GetKeyboardStartY(keyboardType))
	} else {
		menuYAxios = int64(model.GetKeyElementHeight(keyboardType)/2 + model.GetKeyboardStartY(keyboardType))
	}
	return &KeyboardAxiosPos{
		X: menuXAxios,
		Y: menuYAxios,
	}
}

func (model *KeyboardRealConfigModel) RareWordKeyEle(keyboardType KeyboardBrandType) *KeyboardAxiosPos {
	var menuYAxios int64
	var menuXAxios int64
	if keyboardType == HWBaiduBrand {
		menuXAxios = int64(model.GetKeyElementWidth(keyboardType) * 38 / 10)
		menuYAxios = int64(model.GetKeyElementHeight(keyboardType)*12/10 + model.GetKeyboardStartY(keyboardType))
	} else if keyboardType == WeChatBrand {
		menuXAxios = int64(model.GetKeyElementWidth(keyboardType) * 4)
		menuYAxios = int64(model.GetKeyElementHeight(keyboardType)*12/10 + model.GetKeyboardStartY(keyboardType))
	} else {
		menuXAxios = int64(model.GetKeyElementWidth(keyboardType) * 87 / 10)
		menuYAxios = int64(model.GetKeyElementHeight(keyboardType)/2 + model.GetKeyboardStartY(keyboardType))
	}
	return &KeyboardAxiosPos{
		X: menuXAxios,
		Y: menuYAxios,
	}
}

// ChTwentyKeyEle Chinese 26key keyboard element
func (model *KeyboardRealConfigModel) ChTwentyKeyEle(keyboardType KeyboardBrandType) *KeyboardAxiosPos {
	var menuYAxios int64
	var menuXAxios int64
	if keyboardType == HWBaiduBrand {
		menuXAxios = int64(model.GetKeyElementWidth(keyboardType) * 38 / 10)
		menuYAxios = int64(model.GetKeyElementHeight(keyboardType)*12/10 + model.GetKeyboardStartY(keyboardType))
	} else if keyboardType == WeChatBrand {
		menuXAxios = int64(model.GetKeyElementWidth(keyboardType) * 4)
		menuYAxios = int64(model.GetKeyElementHeight(keyboardType)*12/10 + model.GetKeyboardStartY(keyboardType))
	} else {
		menuXAxios = int64(model.GetKeyElementWidth(keyboardType) * 35 / 10)
		menuYAxios = int64(model.GetKeyElementHeight(keyboardType)/2 + model.GetKeyboardStartY(keyboardType))
	}
	return &KeyboardAxiosPos{
		X: menuXAxios,
		Y: menuYAxios,
	}
}

// QwertyQXAxios Get the Q keyboard x axios
func (model *KeyboardRealConfigModel) QwertyQXAxios(keyboardType KeyboardBrandType) int64 {
	var ret int64 = 0
	switch keyboardType {
	case SogouBrand:
		ret = int64(model.BrandDetailModel.Sogou.SlideX + model.GetKeyElementWidth(keyboardType)/2)
	case BaiduBrand:
		ret = int64(model.BrandDetailModel.Baidu.SlideX + model.GetKeyElementWidth(keyboardType)/2)
	case HWBaiduBrand:
		ret = int64(model.BrandDetailModel.HWBaidu.SlideX + model.GetKeyElementWidth(keyboardType)/2)
	case WeChatBrand:
		ret = int64(model.BrandDetailModel.WeChat.SlideX + model.GetKeyElementWidth(keyboardType)/2)
	case IflytekBrand:
		ret = int64(model.BrandDetailModel.Iflytek.SlideX + model.GetKeyElementWidth(keyboardType)/2)
	}
	return ret
}

// QwertyQYAxios Get the Q keyboard y axios
func (model *KeyboardRealConfigModel) QwertyQYAxios(keyboardType KeyboardBrandType) int64 {
	var ret int64 = 0
	ret = int64(model.GetKeyboardStartY(keyboardType) + model.GetKeyElementHeight(keyboardType)/2)
	return ret
}

func (model *KeyboardRealConfigModel) NineSingleItem(keyboardType KeyboardBrandType) uint64 {
	startX := model.GetKeyElementWidth(keyboardType) * 16 / 10
	return startX + model.Nine9KeyWidth(keyboardType)/2
}

// SlideStartY Get the slide input
func (model *KeyboardRealConfigModel) SlideStartY(keyboardType KeyboardBrandType) int64 {
	var ret int64
	if keyboardType == SogouBrand {
		ret = int64(model.SogouStartY + model.BrandDetailModel.Sogou.LogoHeight + model.BrandDetailModel.Sogou.MenuHeight + model.BrandDetailModel.Sogou.SlideY)
	} else if keyboardType == BaiduBrand {
		ret = int64(model.BaiduStartY + model.BrandDetailModel.Baidu.LogoHeight + model.BrandDetailModel.Baidu.MenuHeight + model.BrandDetailModel.Baidu.SlideY)
	} else if keyboardType == IflytekBrand {
		ret = int64(model.IflytekStartY + model.BrandDetailModel.Iflytek.LogoHeight + model.BrandDetailModel.Iflytek.MenuHeight + model.BrandDetailModel.Iflytek.SlideY)
	}
	return ret
}

// SlideStartX Get the x axios of slide x
func (model *KeyboardRealConfigModel) SlideStartX(keyboardType KeyboardBrandType) int64 {
	var ret int64
	if keyboardType == SogouBrand {
		ret = int64(model.BrandDetailModel.Sogou.SlideX)
	} else if keyboardType == BaiduBrand {
		ret = int64(model.BrandDetailModel.Baidu.SlideX)
	} else if keyboardType == IflytekBrand {
		ret = int64(model.BrandDetailModel.Iflytek.SlideX)
	}
	return ret
}

// KRGInstance keyboard shared instance
var KRGInstance *KeyboardRealConfigModel

func KRCGetInstance(deviceName string, filePath string) (single *KeyboardRealConfigModel, keyErr *KeyboardError) {
	if KRGInstance == nil {
		model, err := FileToMobileYaml(filePath)
		if err != nil {
			keyErr = &KeyboardError{
				Message:   "Init mobile keyboard error",
				ErrorType: InitKeyboardError,
			}
			return
		}
		devices := sf.Filter(model.Mobile.BrandNames, func(item BrandNameYamlModel) bool {
			return item.DeviceName == deviceName
		})
		// 2. No devices found in list
		if len(devices) == 0 {
			keyErr = &KeyboardError{
				Message:   "No device found in usb",
				ErrorType: DeviceNotMatchError,
			}
			return
		}
		targetDevice := devices[0]
		targetDevicePos := sf.Filter(model.Mobile.WidthConfig, func(item BrandDetailPosModel) bool {
			return item.IsWidth(targetDevice.Width)
		})[0]
		Once.Do(func() {
			KRGInstance = &KeyboardRealConfigModel{
				SogouStartY:      targetDevice.SogouStartY,
				BaiduStartY:      targetDevice.BaiduStartY,
				IflytekStartY:    targetDevice.IflytekStartY,
				HWBaiduStartY:    targetDevice.HWBaiduStartY,
				WeChatStartY:     targetDevice.WeChatStartY,
				BrandDetailModel: targetDevicePos,
			}
		})
		single = KRGInstance
	}
	return
}

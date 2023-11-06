// Copyright (c) 2011-2023 zhanghao19920218
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package IMEKeyboard

import (
	"sync"
	"time"
)

// KeyboardPos show the keyboard position
type KeyboardPos struct {
	keyboardConfig    KeyboardRealConfigModel
	keyboardBrandType KeyboardBrandType
}

var KBPInstance *KeyboardPos

var Once sync.Once

// MenuKeyboard Switch the keyboard menu
func (model *KeyboardPos) MenuKeyboard() *KeyboardAxiosPos {
	return model.keyboardConfig.MenuPos(model.keyboardBrandType)
}

// DetailSetting Detail Setting of keyboard
func (model *KeyboardPos) DetailSetting() *KeyboardAxiosPos {
	return model.keyboardConfig.DetailMenuPos(model.keyboardBrandType)
}

// NineCh Chinese Nine Key element
func (model *KeyboardPos) NineCh() *KeyboardAxiosPos {
	return model.keyboardConfig.ChNineKeyEle(model.keyboardBrandType)
}

// SixCh
//
//	@Description: 26key keyboard setting
//	@receiver model
//	@return *KeyboardAxiosPos
func (model *KeyboardPos) SixCh() *KeyboardAxiosPos {
	return model.keyboardConfig.ChTwentyKeyEle(model.keyboardBrandType)
}

func (model *KeyboardPos) RareWord() *KeyboardAxiosPos {
	return model.keyboardConfig.RareWordKeyEle(model.keyboardBrandType)
}

func (model *KeyboardPos) KeyboardQ() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

// MoreChooseWord
//
//	@Description: Choose word in more candidate view
//	@receiver model
//	@return *KeyboardAxiosPos
func (model *KeyboardPos) MoreChooseWord() *KeyboardAxiosPos {
	switch model.keyboardBrandType {
	case WeChatBrand:
		return model.KeyboardE()
	default:
		return model.ChooseWord()
	}
}

func (model *KeyboardPos) KeyboardW() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType)),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardE() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*2,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardR() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*3,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardT() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*4,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardY() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*5,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardU() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*6,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardI() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*7,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardO() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*8,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardP() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*9,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardA() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))/2,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardS() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*15/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardD() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*25/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardF() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*35/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardG() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*45/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardH() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*55/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardJ() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*65/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardK() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*75/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardL() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*85/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardZ() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*15/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType))*2,
	}
}

func (model *KeyboardPos) KeyboardX() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*25/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType))*2,
	}
}

func (model *KeyboardPos) KeyboardC() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*35/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType))*2,
	}
}

func (model *KeyboardPos) KeyboardV() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*45/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType))*2,
	}
}

func (model *KeyboardPos) KeyboardB() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*55/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType))*2,
	}
}

func (model *KeyboardPos) KeyboardN() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*65/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType))*2,
	}
}

func (model *KeyboardPos) KeyboardM() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*75/10,
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType))*2,
	}
}

func (model *KeyboardPos) ChooseWord() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.DetailMenuPos(model.keyboardBrandType).X +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType)),
		Y: model.keyboardConfig.DetailMenuPos(model.keyboardBrandType).Y,
	}
}

func (model *KeyboardPos) DeleteAction(keyboardType KeyboardType) *KeyboardAxiosPos {
	var xAxios int64
	var yAxios int64
	xAxios = model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
		int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType)*9)
	if keyboardType == Keyboard26Key {
		yAxios = model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)*2)
	} else {
		yAxios = model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType)
	}
	return &KeyboardAxiosPos{
		X: xAxios,
		Y: yAxios,
	}
}

// KeyboardHeight
//
//	@Description: Get the key element height
//	@receiver model
//	@return uint64
func (model *KeyboardPos) KeyboardHeight() uint64 {
	return model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)
}

// DownArrowAction
//
//	@Description: Get the down arrow action
//	@receiver model
//	@return *KeyboardAxiosPos
func (model *KeyboardPos) DownArrowAction() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType)*9),
		Y: model.keyboardConfig.DetailMenuPos(model.keyboardBrandType).Y,
	}
}

func (model *KeyboardPos) KeyboardOne() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType)),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardABC() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType) +
			model.keyboardConfig.Nine9KeyWidth(model.keyboardBrandType)),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardDEF() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType) +
			model.keyboardConfig.Nine9KeyWidth(model.keyboardBrandType)*2),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType),
	}
}

func (model *KeyboardPos) KeyboardGHI() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType)),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardJKL() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType) +
			model.keyboardConfig.Nine9KeyWidth(model.keyboardBrandType)),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardMNO() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType) +
			model.keyboardConfig.Nine9KeyWidth(model.keyboardBrandType)*2),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)),
	}
}

func (model *KeyboardPos) KeyboardPQRS() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType)),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)*2),
	}
}

func (model *KeyboardPos) KeyboardTUV() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType) +
			model.keyboardConfig.Nine9KeyWidth(model.keyboardBrandType)),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)*2),
	}
}

func (model *KeyboardPos) KeyboardWXYZ() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: int64(model.keyboardConfig.NineSingleItem(model.keyboardBrandType) +
			model.keyboardConfig.Nine9KeyWidth(model.keyboardBrandType)*2),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)*2),
	}
}

func (model *KeyboardPos) KeyboardLastWord() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))*75/10,
		Y: model.keyboardConfig.DetailMenuPos(model.keyboardBrandType).Y,
	}
}

// KeyboardSymbolSwitch
//
//	@Description: Change the keyboard position
//	@receiver model
//	@return *KeyboardAxiosPos
func (model *KeyboardPos) KeyboardSymbolSwitch() *KeyboardAxiosPos {
	return &KeyboardAxiosPos{
		X: model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType),
		Y: model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
			int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType))*3,
	}
}

func (model *KeyboardPos) KeyboardBaiduNextPage() *KeyboardAxiosPos {
	var xAxios int64
	var yAxios = model.keyboardConfig.QwertyQYAxios(model.keyboardBrandType) +
		int64(model.keyboardConfig.GetKeyElementHeight(model.keyboardBrandType)*3)
	if model.keyboardBrandType == HWBaiduBrand {
		xAxios = model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			5*int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))
	} else {
		xAxios = model.keyboardConfig.QwertyQXAxios(model.keyboardBrandType) +
			6*int64(model.keyboardConfig.GetKeyElementWidth(model.keyboardBrandType))
	}
	return &KeyboardAxiosPos{
		X: xAxios,
		Y: yAxios,
	}
}

func (model *KeyboardPos) KeyboardCandidateType() *KeyboardItemSize {
	candidateHeight := model.keyboardConfig.GetCandidateHeight(model.keyboardBrandType)
	imeStart := model.keyboardConfig.ImeStartY(model.keyboardBrandType)
	return &KeyboardItemSize{
		X:      0,
		Y:      int64(imeStart),
		Width:  int64(model.keyboardConfig.BrandDetailModel.Width),
		Height: int64(candidateHeight),
	}
}

var kbpOnce sync.Once

func KBPSharedInstance(
	keyboardBrandType KeyboardBrandType,
	deviceName string,
	filePath string) (single *KeyboardPos, keyErr *KeyboardError) {
	if KBPInstance == nil {
		model, err := KRCGetInstance(deviceName, filePath)
		if err != nil {
			keyErr = &KeyboardError{
				Message:   "Init the keyboard pos error",
				ErrorType: InitKeyboardPosError,
			}
			return
		}
		time.Sleep(1 * time.Second)
		kbpOnce.Do(func() {
			KBPInstance = &KeyboardPos{
				keyboardConfig:    *model,
				keyboardBrandType: keyboardBrandType,
			}
			single = KBPInstance
		})
	}
	return
}

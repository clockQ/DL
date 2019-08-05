/*
 * This file is part of com.pharbers.DL.
 *
 * com.pharbers.DL is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * com.pharbers.DL is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Foobar.  If not, see <https://www.gnu.org/licenses/>.
 */
package PhModel

import (
	"github.com/PharbersDeveloper/DL/PhFactory"
	"reflect"
)

type PhModel struct {
	Model  string
	Query  []map[string]interface{}
	Agg    []map[string]interface{}
	Insert []map[string]interface{}
	Update []map[string]interface{}
	Format []map[string]interface{}
}

func (m PhModel) IsEmpty() bool {
	return reflect.DeepEqual(m, PhModel{})
}

func (m PhModel) FormatResult(data interface{}) (result interface{}, err error) {
	result = data
	table := PhFactory.PhTable{}
	for _, plugin := range m.Format {
		result, err = table.GetFormat(plugin["class"].(string)).Exec(plugin["args"])(data)
		if err != nil {
			return
		}
	}
	return
}

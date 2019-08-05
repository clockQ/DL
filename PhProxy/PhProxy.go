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
package PhProxy

import . "github.com/PharbersDeveloper/DL/PhModel"

type PhProxy interface {
	Create(data PhModel) (err error)
	Update(args PhModel) (data map[string]interface{}, err error)
	Read(args PhModel) (data []map[string]interface{}, err error)
	Delete(args PhModel) (data map[string]interface{}, err error)
}

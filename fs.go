/*
   Copyright (c) 2020 gingfrederik
   Copyright (c) 2021 Gonzalo Fernandez-Victorio
   Copyright (c) 2021 Basement Crowd Ltd (https://www.basementcrowd.com)
   Copyright (c) 2023 Fumiama Minamoto (源文雨)

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package docx

import "embed"

var (
	// TemplateXMLFS stores template docx files
	//go:embed xml
	//go:embed xml/default/_rels/*
	TemplateXMLFS embed.FS

	// DefaultTemplateFilesList is the files list under TemplateXMLFS/xml/default
	DefaultTemplateFilesList = []string{
		"_rels/.rels",
		"docProps/app.xml",
		"docProps/core.xml",
		"word/theme/theme1.xml",
		"word/fontTable.xml",
		"word/styles.xml",
		"[Content_Types].xml",
	}
)

// Copyright © 2017 Ignasi Fosch
//
//  This file is part of pan.
//
//  pan is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Lesser General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  pan is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Lesser General Public License for more details.
//
//  You should have received a copy of the GNU Lesser General Public License
//  along with pan. If not, see <http://www.gnu.org/licenses/>.
//

package pan

import (
	"encoding/xml"
	"testing"
)

var pubDate1 = "Tue, 27 Jan 2015 20:00:00 +0000"
var link1 = "http://mypodcast.com/mypodcast-1.mp3"
var useCases = []struct {
	name    string
	xml     string
	xmlFeed XMLFeed
	ymlFeed YMLFeed
	yml     string
}{
	{
		name: "Empty Rss tag",
		xml: `<?xml version="1.0" encoding="UTF-8"?>
<rss>
</rss>`,
		xmlFeed: XMLFeed{
			XMLName: xml.Name{
				Local: "rss",
			},
			Title: "",
			Items: []XMLItem{},
		},
		ymlFeed: YMLFeed{
			Title: "",
			Items: []YMLItem{},
		},
		yml: "",
	},
	{
		name: "No item list, just channel title",
		xml: `<?xml version="1.0" encoding="UTF-8"?>
<rss>
  <channel>
    <title>My Podcast</title>
  </channel>
</rss>`,
		xmlFeed: XMLFeed{
			XMLName: xml.Name{
				Local: "rss",
			},
			Title: "My Podcast",
			Items: []XMLItem{},
		},
		ymlFeed: YMLFeed{
			Title: "My Podcast",
			Items: []YMLItem{},
		},
		yml: "title: My Podcast",
	},
	{
		name: "Empty item list, just channel title",
		xml: `<?xml version="1.0" encoding="UTF-8"?>
<rss>
  <channel>
    <title>My Podcast</title>
    <item></item>
  </channel>
</rss>`,
		xmlFeed: XMLFeed{
			XMLName: xml.Name{
				Local: "rss",
			},
			Title: "My Podcast",
			Items: []XMLItem{
				{
					XMLName: xml.Name{Local: "item"},
				},
			},
		},
		ymlFeed: YMLFeed{
			Title: "My Podcast",
			Items: []YMLItem{
				{},
			},
		},
		yml: `title: My Podcast
items:
- `,
	},
	{
		name: "Single full item, and channel title",
		xml: `<?xml version="1.0" encoding="UTF-8"?>
<rss>
  <channel>
    <title>My Podcast</title>
    <item>
      <title>My first episode</title>
      <description>Hello world!</description>
      <pubDate>Tue, 27 Jan 2015 20:00:00 +0000</pubDate>
      <link>http://mypodcast.com/mypodcast-1.mp3</link>
    </item>
  </channel>
</rss>`,
		xmlFeed: XMLFeed{
			XMLName: xml.Name{
				Local: "rss",
			},
			Title: "My Podcast",
			Items: []XMLItem{
				{
					XMLName:     xml.Name{Local: "item"},
					Title:       "My first episode",
					Description: "Hello world!",
					PubDate:     pubDate1,
					Link:        link1,
				},
			},
		},
		ymlFeed: YMLFeed{
			Title: "My Podcast",
			Items: []YMLItem{
				{
					Title:       "My first episode",
					Description: "Hello world!",
					PubDate:     pubDate1,
					Link:        link1,
				},
			},
		},
		yml: `title: My Podcast
items:
- title: My first episode
  description: Hello world!
  pubDate: Tue, 27 Jan 2015 20:00:00 +0000
  link: http://mypodcast.com/mypodcast-1.mp3`,
	},
}

func TestReadXML(t *testing.T) {
	for _, useCase := range useCases {
		XMLFeed, err := readXML([]byte(useCase.xml))
		if err != nil {
			t.Errorf("[%s] Unexpected error %s", useCase.name, err)
		}
		if XMLFeed.XMLName != useCase.xmlFeed.XMLName {
			t.Errorf(
				"[%s] Unexpected output xml.Name %s",
				useCase.name,
				XMLFeed.XMLName,
			)
		}
		if XMLFeed.Title != useCase.xmlFeed.Title {
			t.Errorf(
				"[%s] Unexpected output title %s",
				useCase.name,
				XMLFeed.Title,
			)
		}
		if len(useCase.xmlFeed.Items) != len(XMLFeed.Items) {
			t.Errorf(
				"[%s] Wrong number of items %s, should be %s",
				useCase.name,
				len(XMLFeed.Items),
				len(useCase.xmlFeed.Items),
			)
		}
		for i, expectedItem := range useCase.xmlFeed.Items {
			if expectedItem != XMLFeed.Items[i] {
				t.Errorf(
					"[%s] Wrong Item %v",
					useCase.name,
					expectedItem,
				)
			}
		}
	}
}

func TestReadYML(t *testing.T) {
	for _, useCase := range useCases {
		YMLFeed, err := readYML([]byte(useCase.yml))
		if err != nil {
			t.Errorf("[%s] Unexpected error %s", useCase.name, err)
		}
		if YMLFeed.Title != useCase.ymlFeed.Title {
			t.Errorf(
				"[%s] Unexpected output title %s",
				useCase.name,
				YMLFeed.Title,
			)
		}
		if len(useCase.ymlFeed.Items) != len(YMLFeed.Items) {
			t.Errorf(
				"[%s] Wrong number of items %s, should be %s",
				useCase.name,
				len(YMLFeed.Items),
				len(useCase.ymlFeed.Items),
			)
		}
		for i, expectedItem := range useCase.ymlFeed.Items {
			if expectedItem != YMLFeed.Items[i] {
				t.Errorf(
					"[%s] Wrong Item %v",
					useCase.name,
					expectedItem,
				)
			}
		}
	}
}

func TestXML2YML(t *testing.T) {
	for _, useCase := range useCases {
		YMLFeed, err := XML2YML(useCase.xmlFeed)
		if err != nil {
			t.Errorf(
				"[%s] Unexpected error %s",
				useCase.name,
				err,
			)
		}
		if YMLFeed.Title != useCase.ymlFeed.Title {
			t.Errorf(
				"[%s] Unexpected output title %s",
				useCase.name,
				YMLFeed.Title,
			)
		}
		if len(useCase.ymlFeed.Items) != len(YMLFeed.Items) {
			t.Errorf(
				"[%s] Wrong number of items %s, should be %s",
				useCase.name,
				len(YMLFeed.Items),
				len(useCase.ymlFeed.Items),
			)
		}
		for i, expectedItem := range useCase.ymlFeed.Items {
			if expectedItem != YMLFeed.Items[i] {
				t.Errorf(
					"[%s] Wrong Item %v",
					useCase.name,
					expectedItem,
				)
			}
		}
	}
}

func TestYML2XML(t *testing.T) {
	for _, useCase := range useCases {
		XMLFeed, err := YML2XML(useCase.ymlFeed)
		if err != nil {
			t.Errorf(
				"[%s] Unexpected error %s",
				useCase.name,
				err,
			)
		}
		if XMLFeed.XMLName != useCase.xmlFeed.XMLName {
			t.Errorf(
				"[%s] Unexpected output xml.Name %s",
				useCase.name,
				XMLFeed.XMLName,
			)
		}
		if XMLFeed.Title != useCase.xmlFeed.Title {
			t.Errorf(
				"[%s] Unexpected output title %s",
				useCase.name,
				XMLFeed.Title,
			)
		}
		if len(useCase.xmlFeed.Items) != len(XMLFeed.Items) {
			t.Errorf(
				"[%s] Wrong number of items %s, should be %s",
				useCase.name,
				len(XMLFeed.Items),
				len(useCase.xmlFeed.Items),
			)
		}
		for i, expectedItem := range useCase.xmlFeed.Items {
			if expectedItem != XMLFeed.Items[i] {
				t.Errorf(
					"[%s] Wrong ExpectedItem %v",
					useCase.name,
					expectedItem,
				)
			}
		}
	}
}

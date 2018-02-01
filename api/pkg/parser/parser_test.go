package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestCheckForEmptyProps(t *testing.T) {
	tt := []struct {
		name     string
		msg      map[string]string
		expected []string
	}{
		{"Missing Date", map[string]string{
			"Date":       "",
			"From":       "Bravo",
			"Subject":    "Charlie",
			"Message-ID": "Delta",
		}, []string{"Date"}},
		{"Missing From", map[string]string{
			"Date":       "Alpha",
			"From":       "",
			"Subject":    "Charlie",
			"Message-ID": "Delta",
		}, []string{"From"}},
		{"Missing All", map[string]string{
			"Date":       "",
			"From":       "",
			"Subject":    "",
			"Message-ID": "",
		}, []string{"Date", "From", "Subject", "Message-ID"}},
		{"Missing None", map[string]string{
			"Date":       "Alpha",
			"From":       "Bravo",
			"Subject":    "Charlie",
			"Message-ID": "Delta",
		}, []string{}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := checkForEmptyProps(tc.msg)

			if actual != nil {
				for _, prop := range tc.expected {
					if !strings.Contains(actual.Error(), prop) {
						t.Fatalf("checkForEmptyProps of %v\nshould be\n%v\ngot\n%v", tc.name, tc.expected, actual.Error())
					}
				}
			}
		})
	}
}

func TestReadMessageFile(t *testing.T) {
	tt := []struct {
		name string
		file string
	}{}

	_, callerFile, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(callerFile)

	for _, filename := range getTestMessageFilenames("msg", filepath.Join(basePath, "mockMessages")) {
		tt = append(tt, struct {
			name string
			file string
		}{filepath.Base(filename), filename})
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := ReadMessageFile(tc.file)

			b, _ := ioutil.ReadFile(tc.file)

			if bytes.Compare(actual, b) != 0 {
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(string(actual[:]), string(b[:]), true)
				t.Fatalf("ReadMessageFile of %v found a parsing difference of %v", tc.name, dmp.DiffPrettyText(diffs))
			}
		})
	}
}

func TestParseMail(t *testing.T) {
	tt := []struct {
		name     string
		file     string
		expected map[string]string
		err      error
	}{
		{"emptyToHeader.msg", "emptyToHeader.msg",
			map[string]string{
				"Date":       "01 Apr 2011 16:17:41 +0200",
				"From":       "\"Darty\" <infos@contact-darty.com>",
				"Subject":    "Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
				"Message-ID": "<20110401161739.E3786358A9D7B977@contact-darty.com>",
			},
			fmt.Errorf("missing or malformed email header(s): To"),
		},
		{"emptyDateHeader.msg", "emptyDateHeader.msg",
			map[string]string{
				"From":       "\"Darty\" <infos@contact-darty.com>",
				"To":         "1000mercis@cp.assurance.returnpath.net",
				"Subject":    "Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
				"Message-ID": "<20110401161739.E3786358A9D7B977@contact-darty.com>",
			},
			fmt.Errorf("missing or malformed email header(s): Date"),
		},
		{"emptyFromHeader.msg", "emptyFromHeader.msg",
			map[string]string{
				"Date":       "01 Apr 2011 16:17:41 +0200",
				"To":         "1000mercis@cp.assurance.returnpath.net",
				"Subject":    "Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
				"Message-ID": "<20110401161739.E3786358A9D7B977@contact-darty.com>",
			},
			fmt.Errorf("missing or malformed email header(s): From"),
		},
		{"emptyMessageIdHeader.msg", "emptyMessageIdHeader.msg",
			map[string]string{
				"Date":    "01 Apr 2011 16:17:41 +0200",
				"From":    "\"Darty\" <infos@contact-darty.com>",
				"To":      "1000mercis@cp.assurance.returnpath.net",
				"Subject": "Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
			},
			fmt.Errorf("missing or malformed email header(s): Message-ID"),
		},
		{"emptySubjectHeader.msg", "emptySubjectHeader.msg",
			map[string]string{
				"Date":       "01 Apr 2011 16:17:41 +0200",
				"From":       "\"Darty\" <infos@contact-darty.com>",
				"To":         "1000mercis@cp.assurance.returnpath.net",
				"Message-ID": "<20110401161739.E3786358A9D7B977@contact-darty.com>",
			},
			fmt.Errorf("missing or malformed email header(s): Subject"),
		},
		{"emptyContentTypeHeader.msg", "emptyContentTypeHeader.msg",
			map[string]string{
				"Date":       "Fri, 01 Apr 2011 05:52:55 PDT",
				"From":       "Corel <news@email1-corel.com>",
				"To":         "corel@cp.monitor1.returnpath.net",
				"Subject":    "PREVIEW:   Save $170 and get special gift with CorelDraw Premium Suite X5",
				"Message-ID": "<Corel.6k3yh-636g-.fv4t@email1-corel.com>",
			},
			fmt.Errorf("missing or malformed email header(s): Content-Type"),
		},
		{"20110401_1000mercis_14461469_html.msg", "20110401_1000mercis_14461469_html.msg",
			map[string]string{
				"Date":         "01 Apr 2011 16:17:41 +0200",
				"From":         "\"Darty\" <infos@contact-darty.com>",
				"To":           "1000mercis@cp.assurance.returnpath.net",
				"Subject":      "Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
				"Message-ID":   "<20110401161739.E3786358A9D7B977@contact-darty.com>",
				"Content-Type": "text/html; charset=\"iso-8859-1\"",
			},
			nil,
		},
		{"20110401_aamarketinginc_14456749_html.msg", "20110401_aamarketinginc_14456749_html.msg",
			map[string]string{
				"Date":         "Thu, 31 Mar 2011 23:19:52 -0500",
				"From":         "MindsPay<survey@mindspaymails.com>",
				"To":           "aamarketinginc@cp.monitor1.returnpath.net",
				"Subject":      "Paid Mail : Offer #10491 get $4.00",
				"Message-ID":   "<MP1301631592801EH10491@mindspay.com>",
				"Content-Type": "text/html; charset=iso-8859-1",
			},
			nil,
		},
		{"20110401_aeg_14465739_html.msg", "20110401_aeg_14465739_html.msg",
			map[string]string{
				"Date":         "Fri, 01 Apr 2011 12:06:22 -0600",
				"From":         "\"LA Galaxy\" <enews@events.lagalaxy.com>",
				"To":           "<aeg@cp.delivery.ncrcustomerpower.com>",
				"Subject":      "April Fool's Day Offer, Save up to 40% with no fees",
				"Message-ID":   "<3709e1a3-663f-464c-a38f-584ae8c9fe24@xtinmta105.xt.local>",
				"Content-Type": "text/html; charset=\"us-ascii\"",
			},
			nil},
		{"20110401_alchemyworx_14461429_multialt.msg", "20110401_alchemyworx_14461429_multialt.msg",
			map[string]string{
				"Date":         "Fri, 1 Apr 2011 14:14:49 -0000",
				"From":         "=?utf-8?q?Aviva?= <aviva@avivaemail.co.uk>",
				"To":           "alchemyworx@cp.assurance.returnpath.net",
				"Subject":      "(TEST-Multipart) =?utf-8?q?=5BRetention_In_Life_ezine=5Fhome=5F050411=5D_Introducing_Your_?= =?utf-8?q?Aviva_Essentials=3A_Win_4_tickets_to_the_Aviva_Premiership_Rugb?= =?utf-8?q?y_Final=2C_Keep_the_cost_of_driving_down_and_more?=",
				"Message-ID":   "<bx6rw3raupta74au6m0rxbysph09qe.0.15@mta141.avivaemail.co.uk>",
				"Content-Type": "multipart/alternative; boundary=\"=bx6rw3raupta74au6m0rxbysph09qe\"",
			}, nil},
		{"20110401_americancollegiatemarketing_14461959_multialt.msg", "20110401_americancollegiatemarketing_14461959_multialt.msg",
			map[string]string{
				"Date":         "Fri, 1 Apr 2011 10:38:16 -0400",
				"From":         "<Amway@MagazineLine.com>",
				"To":           "<americancollegiatemarketing@cp.monitor1.returnpath.net>",
				"Subject":      "April 2011 TPM Amway No Subs Spring Savings from MagazineLine",
				"Message-ID":   "<EF9C090C1310457C97AD9E1279F0BF68@acm.local>",
				"Content-Type": "multipart/alternative; boundary=\"----=_NextPart_000_0032_01CBF058.E8C87270\"",
			}, nil},
		{"20110401_beliefnet_14461399_html.msg", "20110401_beliefnet_14461399_html.msg",
			map[string]string{
				"Date":         "Fri,  1 Apr 2011 08:12:00 -0600 (MDT)",
				"From":         "Announce - Beliefnet Sponsor <specialoffers@mail.beliefnet.com>",
				"To":           "<beliefnet@cp.monitor1.returnpath.net>",
				"Subject":      "[SP] Grant Funding May Be Available for Top Online Colleges. Get Free Info Today.",
				"Message-ID":   "<527817310.344.1301667087687.JavaMail.root@mail.beliefnet.com>",
				"Content-Type": "text/html; charset=UTF-8",
			}, nil},
		{"20110401_beliefnet_14464159_html.msg", "20110401_beliefnet_14464159_html.msg",
			map[string]string{
				"Date":         "Fri,  1 Apr 2011 10:32:42 -0600 (MDT)",
				"From":         "Chicken Soup - Beliefnet Partner <specialoffers@mail.beliefnet.com>",
				"To":           "<beliefnet@cp.monitor1.returnpath.net>",
				"Subject":      "[SP] The Art of Positive Thinking",
				"Message-ID":   "<463918295.411.1301674909118.JavaMail.root@mail.beliefnet.com>",
				"Content-Type": "text/html; charset=UTF-8",
			}, nil},
		{"20110401_boydgamingcorporation_14465279_multialt.msg", "20110401_boydgamingcorporation_14465279_multialt.msg",
			map[string]string{
				"Date":         "Fri, 01 Apr 2011 10:36:26 -0700",
				"From":         "Suncoast Hotel & Casino - Las Vegas <suncoast@boydgaming.net>",
				"To":           "Lisa Marshall <boydgamingcorporation@cp.assurance.returnpath.net>",
				"Subject":      "See What's Happening with our Table Games!",
				"Message-ID":   "<20110401173626.15575.2089030531.swift@webadmin.boydgaming.net>",
				"Content-Type": "multipart/alternative; boundary=\"_=_swift-13596511954d960d1a312b33.37345773_=_\"",
			}, nil},
		{"20110401_citibanksingaporelimited_14456499_multialt.msg", "20110401_citibanksingaporelimited_14456499_multialt.msg",
			map[string]string{
				"Date":         "Fri, 1 Apr 2011 02:57:21 GMT",
				"From":         "customer.service@citicorp.com",
				"To":           "citibanksingaporelimited@cp.monitor1.returnpath.net",
				"Subject":      "Citi Alerts",
				"Message-ID":   "<1479419471.1301626641534.JavaMail.pjfpbg1@saixp36>",
				"Content-Type": "multipart/alternative; boundary=\"----=_Part_26607_1527703119.1301626641533\"",
			}, nil},
		{"20110401_cobaltgroup_14464029_html.msg", "20110401_cobaltgroup_14464029_html.msg",
			map[string]string{
				"Date":         "Fri, 1 Apr 2011 16:25:17 +0000 (UTC)",
				"From":         "Hometown Motors <klongfield.10162425@dealer.onstation.com>",
				"To":           "cobaltgroup@cp.monitor1.returnpath.net",
				"Subject":      "[SAMPLE] 04-719314-2011 Chevy April DAP #1",
				"Message-ID":   "<6426946.1413.1301675117949.JavaMail.tomcat@osadmin02>",
				"Content-Type": "text/html; charset=UTF-8",
			}, nil},
		{"20110401_compostmarketingab_14459379_multialt.msg", "20110401_compostmarketingab_14459379_multialt.msg",
			map[string]string{
				"Date":         "Fri, 1 Apr 2011 13:02:13 +0200",
				"From":         "{CARMA TEST} Test <from@test.carmamail.com>",
				"To":           "compostmarketingab@cp.monitor1.returnpath.net",
				"Subject":      "ComHem Senaste Nyheterna",
				"Message-ID":   "<58795828.1301655732499.JavaMail.compostadmin@secos-a107>",
				"Content-Type": "multipart/alternative; boundary=\"----=_Part_329666_58751540.1301655732499\"",
			}, nil},
		{"20110401_corel_14460139_html.msg", "20110401_corel_14460139_html.msg",
			map[string]string{
				"Date":         "Fri, 01 Apr 2011 05:52:55 PDT",
				"From":         "Corel <news@email1-corel.com>",
				"To":           "corel@cp.monitor1.returnpath.net",
				"Subject":      "PREVIEW:   Save $170 and get special gift with CorelDraw Premium Suite X5",
				"Message-ID":   "<Corel.6k3yh-636g-.fv4t@email1-corel.com>",
				"Content-Type": "text/html; charset=us-ascii",
			}, nil},
	}

	_, callerFile, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(callerFile)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			msgBytes, _ := ReadMessageFile(filepath.Join(basePath, "mockMessages", tc.file))
			msg := ParseMail(string(msgBytes[:]))

			actual, GetMessageErr := GetMessage(msg, "Date", "From", "To", "Subject", "Message-ID", "Content-Type")

			if GetMessageErr == nil && !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("ParseMail of %v \nshould be:\n; %v\n; got:\n; %v\n;", tc.name, tc.expected, actual)
			}

			if (GetMessageErr != nil) && GetMessageErr.Error() != tc.err.Error() {
				t.Fatalf("ParseMail of %v\n should've thrown error: %v\n got: %v", tc.name, tc.err, GetMessageErr)
			}
		})
	}
}

func getTestMessageFilenames(ext, dir string) []string {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	messageFiles := []string{}
	for _, file := range files {
		if ok, _ := regexp.Match(ext, []byte(file.Name())); ok {
			messageFiles = append(messageFiles, filepath.Join(dir, file.Name()))
		}
	}

	return messageFiles
}

//go:build !custom || parsers || parsers.avro

package all

import _ "github.com/influxdata/telegraf/plugins/parsers/avro" // register plugin

// Code generated by qtc from "debug.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line lib/promrelabel/debug.qtpl:1
package promrelabel

//line lib/promrelabel/debug.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/htmlcomponents"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promutils"
)

//line lib/promrelabel/debug.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/promrelabel/debug.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/promrelabel/debug.qtpl:8
func StreamRelabelDebugSteps(qw422016 *qt422016.Writer, isTargetRelabel bool, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) {
//line lib/promrelabel/debug.qtpl:8
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head>`)
//line lib/promrelabel/debug.qtpl:12
	htmlcomponents.StreamCommonHeader(qw422016)
//line lib/promrelabel/debug.qtpl:12
	qw422016.N().S(`<title>Metric relabel debug</title><script>function submitRelabelDebugForm(e) {var form = e.target;var method = "GET";if (form.elements["relabel_configs"].value.length + form.elements["metric"].value.length > 1000) {method = "POST";}form.method = method;}</script></head><body>`)
//line lib/promrelabel/debug.qtpl:26
	htmlcomponents.StreamNavbar(qw422016)
//line lib/promrelabel/debug.qtpl:26
	qw422016.N().S(`<div class="container-fluid"><a href="https://docs.victoriametrics.com/relabeling.html" target="_blank">Relabeling docs</a>`)
//line lib/promrelabel/debug.qtpl:28
	qw422016.N().S(` `)
//line lib/promrelabel/debug.qtpl:30
	if isTargetRelabel {
//line lib/promrelabel/debug.qtpl:30
		qw422016.N().S(`<a href="metric-relabel-debug`)
//line lib/promrelabel/debug.qtpl:31
		if targetID != "" {
//line lib/promrelabel/debug.qtpl:31
			qw422016.N().S(`?id=`)
//line lib/promrelabel/debug.qtpl:31
			qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:31
		}
//line lib/promrelabel/debug.qtpl:31
		qw422016.N().S(`">Metric relabel debug</a>`)
//line lib/promrelabel/debug.qtpl:32
	} else {
//line lib/promrelabel/debug.qtpl:32
		qw422016.N().S(`<a href="target-relabel-debug`)
//line lib/promrelabel/debug.qtpl:33
		if targetID != "" {
//line lib/promrelabel/debug.qtpl:33
			qw422016.N().S(`?id=`)
//line lib/promrelabel/debug.qtpl:33
			qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:33
		}
//line lib/promrelabel/debug.qtpl:33
		qw422016.N().S(`">Target relabel debug</a>`)
//line lib/promrelabel/debug.qtpl:34
	}
//line lib/promrelabel/debug.qtpl:34
	qw422016.N().S(`<br>`)
//line lib/promrelabel/debug.qtpl:37
	if err != nil {
//line lib/promrelabel/debug.qtpl:38
		htmlcomponents.StreamErrorNotification(qw422016, err)
//line lib/promrelabel/debug.qtpl:39
	}
//line lib/promrelabel/debug.qtpl:39
	qw422016.N().S(`<div class="m-3"><form method="POST" onsubmit="submitRelabelDebugForm(event)">`)
//line lib/promrelabel/debug.qtpl:43
	streamrelabelDebugFormInputs(qw422016, metric, relabelConfigs)
//line lib/promrelabel/debug.qtpl:44
	if targetID != "" {
//line lib/promrelabel/debug.qtpl:44
		qw422016.N().S(`<input type="hidden" name="id" value="`)
//line lib/promrelabel/debug.qtpl:45
		qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:45
		qw422016.N().S(`" />`)
//line lib/promrelabel/debug.qtpl:46
	}
//line lib/promrelabel/debug.qtpl:46
	qw422016.N().S(`<input type="submit" value="Submit" class="btn btn-primary m-1" />`)
//line lib/promrelabel/debug.qtpl:48
	if targetID != "" {
//line lib/promrelabel/debug.qtpl:48
		qw422016.N().S(`<button type="button" onclick="location.href='?id=`)
//line lib/promrelabel/debug.qtpl:49
		qw422016.E().S(targetID)
//line lib/promrelabel/debug.qtpl:49
		qw422016.N().S(`'" class="btn btn-secondary m-1">Reset</button>`)
//line lib/promrelabel/debug.qtpl:50
	}
//line lib/promrelabel/debug.qtpl:50
	qw422016.N().S(`</form></div><div class="row"><main class="col-12">`)
//line lib/promrelabel/debug.qtpl:56
	streamrelabelDebugSteps(qw422016, dss)
//line lib/promrelabel/debug.qtpl:56
	qw422016.N().S(`</main></div></div></body></html>`)
//line lib/promrelabel/debug.qtpl:62
}

//line lib/promrelabel/debug.qtpl:62
func WriteRelabelDebugSteps(qq422016 qtio422016.Writer, isTargetRelabel bool, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) {
//line lib/promrelabel/debug.qtpl:62
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:62
	StreamRelabelDebugSteps(qw422016, isTargetRelabel, targetID, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:62
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:62
}

//line lib/promrelabel/debug.qtpl:62
func RelabelDebugSteps(isTargetRelabel bool, targetID string, dss []DebugStep, metric, relabelConfigs string, err error) string {
//line lib/promrelabel/debug.qtpl:62
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:62
	WriteRelabelDebugSteps(qb422016, isTargetRelabel, targetID, dss, metric, relabelConfigs, err)
//line lib/promrelabel/debug.qtpl:62
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:62
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:62
	return qs422016
//line lib/promrelabel/debug.qtpl:62
}

//line lib/promrelabel/debug.qtpl:64
func streamrelabelDebugFormInputs(qw422016 *qt422016.Writer, metric, relabelConfigs string) {
//line lib/promrelabel/debug.qtpl:64
	qw422016.N().S(`<div>Relabel configs:<br/><textarea name="relabel_configs" style="width: 100%; height: 15em" class="m-1">`)
//line lib/promrelabel/debug.qtpl:67
	qw422016.E().S(relabelConfigs)
//line lib/promrelabel/debug.qtpl:67
	qw422016.N().S(`</textarea></div><div>Labels:<br/><textarea name="metric" style="width: 100%; height: 5em" class="m-1">`)
//line lib/promrelabel/debug.qtpl:72
	qw422016.E().S(metric)
//line lib/promrelabel/debug.qtpl:72
	qw422016.N().S(`</textarea></div>`)
//line lib/promrelabel/debug.qtpl:74
}

//line lib/promrelabel/debug.qtpl:74
func writerelabelDebugFormInputs(qq422016 qtio422016.Writer, metric, relabelConfigs string) {
//line lib/promrelabel/debug.qtpl:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:74
	streamrelabelDebugFormInputs(qw422016, metric, relabelConfigs)
//line lib/promrelabel/debug.qtpl:74
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:74
}

//line lib/promrelabel/debug.qtpl:74
func relabelDebugFormInputs(metric, relabelConfigs string) string {
//line lib/promrelabel/debug.qtpl:74
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:74
	writerelabelDebugFormInputs(qb422016, metric, relabelConfigs)
//line lib/promrelabel/debug.qtpl:74
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:74
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:74
	return qs422016
//line lib/promrelabel/debug.qtpl:74
}

//line lib/promrelabel/debug.qtpl:76
func streamrelabelDebugSteps(qw422016 *qt422016.Writer, dss []DebugStep) {
//line lib/promrelabel/debug.qtpl:77
	if len(dss) > 0 {
//line lib/promrelabel/debug.qtpl:77
		qw422016.N().S(`<div class="m-3"><b>Original labels:</b> <samp>`)
//line lib/promrelabel/debug.qtpl:79
		streammustFormatLabels(qw422016, dss[0].In)
//line lib/promrelabel/debug.qtpl:79
		qw422016.N().S(`</samp></div>`)
//line lib/promrelabel/debug.qtpl:81
	}
//line lib/promrelabel/debug.qtpl:81
	qw422016.N().S(`<table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col" style="width: 5%">Step</th><th scope="col" style="width: 25%">Relabeling Rule</th><th scope="col" style="width: 35%">Input Labels</th><th scope="col" stile="width: 35%">Output labels</a></tr></thead><tbody>`)
//line lib/promrelabel/debug.qtpl:92
	for i, ds := range dss {
//line lib/promrelabel/debug.qtpl:94
		inLabels := promutils.MustNewLabelsFromString(ds.In)
		outLabels := promutils.MustNewLabelsFromString(ds.Out)
		changedLabels := getChangedLabelNames(inLabels, outLabels)

//line lib/promrelabel/debug.qtpl:97
		qw422016.N().S(`<tr><td>`)
//line lib/promrelabel/debug.qtpl:99
		qw422016.N().D(i)
//line lib/promrelabel/debug.qtpl:99
		qw422016.N().S(`</td><td><b><pre class="m-2">`)
//line lib/promrelabel/debug.qtpl:100
		qw422016.E().S(ds.Rule)
//line lib/promrelabel/debug.qtpl:100
		qw422016.N().S(`</pre></b></td><td><div class="m-2" style="font-size: 0.9em" title="deleted and updated labels highlighted in red">`)
//line lib/promrelabel/debug.qtpl:103
		streamlabelsWithHighlight(qw422016, inLabels, changedLabels, "red")
//line lib/promrelabel/debug.qtpl:103
		qw422016.N().S(`</div></td><td><div class="m-2" style="font-size: 0.9em" title="added and updated labels highlighted in blue">`)
//line lib/promrelabel/debug.qtpl:108
		streamlabelsWithHighlight(qw422016, outLabels, changedLabels, "blue")
//line lib/promrelabel/debug.qtpl:108
		qw422016.N().S(`</div></td></tr>`)
//line lib/promrelabel/debug.qtpl:112
	}
//line lib/promrelabel/debug.qtpl:112
	qw422016.N().S(`</tbody></table>`)
//line lib/promrelabel/debug.qtpl:115
	if len(dss) > 0 {
//line lib/promrelabel/debug.qtpl:115
		qw422016.N().S(`<div class="m-3"><b>Resulting labels:</b> <samp>`)
//line lib/promrelabel/debug.qtpl:117
		streammustFormatLabels(qw422016, dss[len(dss)-1].Out)
//line lib/promrelabel/debug.qtpl:117
		qw422016.N().S(`</samp></div>`)
//line lib/promrelabel/debug.qtpl:119
	}
//line lib/promrelabel/debug.qtpl:120
}

//line lib/promrelabel/debug.qtpl:120
func writerelabelDebugSteps(qq422016 qtio422016.Writer, dss []DebugStep) {
//line lib/promrelabel/debug.qtpl:120
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:120
	streamrelabelDebugSteps(qw422016, dss)
//line lib/promrelabel/debug.qtpl:120
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:120
}

//line lib/promrelabel/debug.qtpl:120
func relabelDebugSteps(dss []DebugStep) string {
//line lib/promrelabel/debug.qtpl:120
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:120
	writerelabelDebugSteps(qb422016, dss)
//line lib/promrelabel/debug.qtpl:120
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:120
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:120
	return qs422016
//line lib/promrelabel/debug.qtpl:120
}

//line lib/promrelabel/debug.qtpl:122
func streamlabelsWithHighlight(qw422016 *qt422016.Writer, labels *promutils.Labels, highlight map[string]struct{}, color string) {
//line lib/promrelabel/debug.qtpl:124
	labelsList := labels.GetLabels()
	metricName := ""
	for i, label := range labelsList {
		if label.Name == "__name__" {
			metricName = label.Value
			labelsList = append(labelsList[:i], labelsList[i+1:]...)
			break
		}
	}

//line lib/promrelabel/debug.qtpl:134
	if metricName != "" {
//line lib/promrelabel/debug.qtpl:135
		if _, ok := highlight["__name__"]; ok {
//line lib/promrelabel/debug.qtpl:135
			qw422016.N().S(`<span style="font-weight:bold;color:`)
//line lib/promrelabel/debug.qtpl:136
			qw422016.E().S(color)
//line lib/promrelabel/debug.qtpl:136
			qw422016.N().S(`">`)
//line lib/promrelabel/debug.qtpl:136
			qw422016.E().S(metricName)
//line lib/promrelabel/debug.qtpl:136
			qw422016.N().S(`</span>`)
//line lib/promrelabel/debug.qtpl:137
		} else {
//line lib/promrelabel/debug.qtpl:138
			qw422016.E().S(metricName)
//line lib/promrelabel/debug.qtpl:139
		}
//line lib/promrelabel/debug.qtpl:140
		if len(labelsList) == 0 {
//line lib/promrelabel/debug.qtpl:140
			return
//line lib/promrelabel/debug.qtpl:140
		}
//line lib/promrelabel/debug.qtpl:141
	}
//line lib/promrelabel/debug.qtpl:141
	qw422016.N().S(`{`)
//line lib/promrelabel/debug.qtpl:143
	for i, label := range labelsList {
//line lib/promrelabel/debug.qtpl:144
		if _, ok := highlight[label.Name]; ok {
//line lib/promrelabel/debug.qtpl:144
			qw422016.N().S(`<span style="font-weight:bold;color:`)
//line lib/promrelabel/debug.qtpl:145
			qw422016.E().S(color)
//line lib/promrelabel/debug.qtpl:145
			qw422016.N().S(`">`)
//line lib/promrelabel/debug.qtpl:145
			qw422016.E().S(label.Name)
//line lib/promrelabel/debug.qtpl:145
			qw422016.N().S(`=`)
//line lib/promrelabel/debug.qtpl:145
			qw422016.E().Q(label.Value)
//line lib/promrelabel/debug.qtpl:145
			qw422016.N().S(`</span>`)
//line lib/promrelabel/debug.qtpl:146
		} else {
//line lib/promrelabel/debug.qtpl:147
			qw422016.E().S(label.Name)
//line lib/promrelabel/debug.qtpl:147
			qw422016.N().S(`=`)
//line lib/promrelabel/debug.qtpl:147
			qw422016.E().Q(label.Value)
//line lib/promrelabel/debug.qtpl:148
		}
//line lib/promrelabel/debug.qtpl:149
		if i < len(labelsList)-1 {
//line lib/promrelabel/debug.qtpl:149
			qw422016.N().S(`,`)
//line lib/promrelabel/debug.qtpl:149
			qw422016.N().S(` `)
//line lib/promrelabel/debug.qtpl:149
		}
//line lib/promrelabel/debug.qtpl:150
	}
//line lib/promrelabel/debug.qtpl:150
	qw422016.N().S(`}`)
//line lib/promrelabel/debug.qtpl:152
}

//line lib/promrelabel/debug.qtpl:152
func writelabelsWithHighlight(qq422016 qtio422016.Writer, labels *promutils.Labels, highlight map[string]struct{}, color string) {
//line lib/promrelabel/debug.qtpl:152
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:152
	streamlabelsWithHighlight(qw422016, labels, highlight, color)
//line lib/promrelabel/debug.qtpl:152
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:152
}

//line lib/promrelabel/debug.qtpl:152
func labelsWithHighlight(labels *promutils.Labels, highlight map[string]struct{}, color string) string {
//line lib/promrelabel/debug.qtpl:152
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:152
	writelabelsWithHighlight(qb422016, labels, highlight, color)
//line lib/promrelabel/debug.qtpl:152
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:152
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:152
	return qs422016
//line lib/promrelabel/debug.qtpl:152
}

//line lib/promrelabel/debug.qtpl:154
func streammustFormatLabels(qw422016 *qt422016.Writer, s string) {
//line lib/promrelabel/debug.qtpl:155
	labels := promutils.MustNewLabelsFromString(s)

//line lib/promrelabel/debug.qtpl:156
	streamlabelsWithHighlight(qw422016, labels, nil, "")
//line lib/promrelabel/debug.qtpl:157
}

//line lib/promrelabel/debug.qtpl:157
func writemustFormatLabels(qq422016 qtio422016.Writer, s string) {
//line lib/promrelabel/debug.qtpl:157
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promrelabel/debug.qtpl:157
	streammustFormatLabels(qw422016, s)
//line lib/promrelabel/debug.qtpl:157
	qt422016.ReleaseWriter(qw422016)
//line lib/promrelabel/debug.qtpl:157
}

//line lib/promrelabel/debug.qtpl:157
func mustFormatLabels(s string) string {
//line lib/promrelabel/debug.qtpl:157
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promrelabel/debug.qtpl:157
	writemustFormatLabels(qb422016, s)
//line lib/promrelabel/debug.qtpl:157
	qs422016 := string(qb422016.B)
//line lib/promrelabel/debug.qtpl:157
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promrelabel/debug.qtpl:157
	return qs422016
//line lib/promrelabel/debug.qtpl:157
}

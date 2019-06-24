
<template>
	<div id = "test">
		<div id="block"></div>
		<iframe src="../../static/html/main.html" id = "frame1"/>
	</div>
</template>

<script>

export default {
  name: 'App',
	data() {
      return {
				rawHtml : 'hello world'
		}
	}
}
</script>

<style>
#frame1 {
	position: absolute;
	top: 10%;
	height: 90%;
	width: 100%;
}
</style>

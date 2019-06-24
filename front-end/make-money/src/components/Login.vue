
<template>
	<div id = "test">
		<div id="block"></div>
		<iframe src="../../static/html/login.html" id = "frame1"/>
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
#block {
	height: 8rem;;
}
#frame1 {
	position: absolute;
	height: 90%;
	width: 100%;
}
</style>


<template>
	<div id = "test">
		<iframe src="./static/html/login.html" id = "frame1"/>
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

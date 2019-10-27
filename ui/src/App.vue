<template>
	<div id="app" class="md-app md-layout">
		<md-toolbar>
			<div class="md-title">dream01</div>
		</md-toolbar>
		<md-content class="md-size-100">
			<router-view />
			<player></player>
		</md-content>
	</div>
</template>

<script>
import Player from "./components/Player.vue";

export default {
	components: {
		Player
	}
};
</script>

<style lang="scss">
@font-face {
	font-family: "strange";
	src: url("./assets/fonts/Benguiat-Bold.ttf");
}

@font-face {
	font-family: "mrrobot";
	src: url("./assets/fonts/mr-robot.ttf");
}

.md-title {
	font-family: "mrrobot";
}
</style>

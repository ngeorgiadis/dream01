<template>
	<div id="app">
		<md-app md-waterfall md-mode="fixed" md-tag="div">
			<md-app-toolbar class="md-theme-extra md-primary">
				<div class="main-toolbar md-title">dream01</div>
			</md-app-toolbar>

			<md-app-content class="main-content">
				<router-view></router-view>
			</md-app-content>

			<!-- <md-app-toolbar class="player-bar md-primary md-theme-footer">
				<div class="  md-title">dream01</div>
			</md-app-toolbar> -->
			<md-app-content class="footer md-theme-footer md-accent">
				<radio-player></radio-player>
			</md-app-content>
		</md-app>
	</div>
</template>

<script>
import RadioPlayer from "./components/Player.vue";

export default {
	components: {
		RadioPlayer
	},

	data() {
		return {
			menuVisible: false
		};
	},

	methods: {
		toggleMenu() {
			this.menuVisible = !this.menuVisible;
		}
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

@import "~vue-material/dist/theme/engine";
@include md-register-theme(
	"extra",
	(
		primary: md-get-palette-color(blue, 900),
		accent: md-get-palette-color(red, A200)
	)
);
@import "~vue-material/dist/components/MdToolbar/theme";

@include md-register-theme(
	"footer",
	(
		primary: md-get-palette-color(white, 900),
		accent: md-get-palette-color(red, 900)
	)
);
@import "~vue-material/dist/components/MdContent/theme";
@import "~vue-material/dist/components/MdButton/theme";

//@import "~vue-material/dist/theme/all";

.md-app {
	height: 100vh;
	//border: 1px solid rgba(#000, 0.12);
}

.md-drawer {
	width: 230px;
	max-width: calc(100vw - 125px);
}

.logo img {
	height: 40px;
}

.router-link-active {
	background-color: #eee;
}

.main-toolbar {
	font-family: "mrrobot";
}

.player-bar {
	font-family: "strange";
}

.md-app-scroller {
	margin-bottom: 64px !important;
}

.footer {
	height: 64px;

	bottom: 0;
	left: 0;
	right: 0;

	position: fixed;
	// z-index: 100;
}
</style>

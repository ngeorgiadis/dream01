<template>
	<v-app>
		<div id="main">
			<v-app-bar height="70px" dark app class="main-toolbar">
				<router-link class="no-background" to="/">
					<img
						class="logo"
						src="./assets/logo-small-clean.png"
						alt=""
					/>
				</router-link>

				<v-toolbar>
					<div>dream01</div>
					<v-spacer></v-spacer>

					<v-btn to="/discover" icon>
						<v-icon>mdi-magnify</v-icon>
					</v-btn>
				</v-toolbar>
			</v-app-bar>

			<v-content class="main-content">
				<router-view></router-view>
			</v-content>

			<v-footer fixed dark class="footer">
				<radio-player></radio-player>
			</v-footer>
		</div>
	</v-app>
</template>

<script>
import RadioPlayer from "./components/Player.vue";

export default {
	components: {
		RadioPlayer,
	},

	data() {
		return {
			menuVisible: false,
		};
	},

	methods: {
		toggleMenu() {
			this.menuVisible = !this.menuVisible;
		},
	},
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

.router-link-active {
	background-color: #eee;
}

.main-toolbar {
	font-family: "mrrobot";
}

.main-toolbar div {
	font-size: 34px;
	font-weight: normal;
	color: #eeeeee;
}

.player-bar {
	font-family: "strange";
}

.main-content {
	overflow: auto;
	margin-bottom: 100px;
}

.logo {
	margin-top: -30px;
}

.no-background {
	background-color: unset;
}
</style>

export type RootStackParamList = {
    SplashScreen: {}
    Register: {}
    Login: {}
    // Home: undefined;
    Home: { username: string };
    DrawNavigationRoutes: {}
    // Feed: { sort: 'latest' | 'top' } | undefined;
};
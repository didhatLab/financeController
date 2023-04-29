import {StyleSheet, Text, View} from 'react-native';
import * as React from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import {createDrawerNavigator} from '@react-navigation/drawer';
import {RootStackParamList} from "./types/screen-types";
import {HomeScreen} from "./screens/menu/home";
import {Provider as PaperProvider} from 'react-native-paper';
import {SplashScreen} from "./screens/splash";
import {DrawerNavigatorRoutes} from "./screens/menu/drawer-navigator-routes";
import {LoginScreen} from "./screens/auth/login";
import RegisterScreen from "./screens/auth/register";
import {NativeBaseProvider} from "native-base";

const Stack = createNativeStackNavigator<RootStackParamList>()
const Drawer = createDrawerNavigator()

export default function App() {
    return (
        <NativeBaseProvider>
            <PaperProvider>
                <NavigationContainer>
                    <Stack.Navigator initialRouteName={"SplashScreen"}>
                        <Stack.Screen name={"SplashScreen"} component={SplashScreen} options={{}}/>
                        <Stack.Screen name={"DrawNavigationRoutes"} component={DrawerNavigatorRoutes}/>
                        <Stack.Screen name={"Login"} component={LoginScreen}/>
                        <Stack.Screen name={"Register"} component={RegisterScreen}/>
                    </Stack.Navigator>
                </NavigationContainer>
            </PaperProvider>
        </NativeBaseProvider>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
    },
});

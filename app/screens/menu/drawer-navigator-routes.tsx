import {createStackNavigator} from '@react-navigation/stack';
import {createDrawerNavigator} from '@react-navigation/drawer';

import {RootStackParamList} from "../../types/screen-types";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {useEffect, useState} from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {HomeScreen} from "./home";
import {DrawerParams} from "./types";
import {AddNewSpendScreen} from "./add-spending";

type Props = NativeStackScreenProps<RootStackParamList, 'DrawNavigationRoutes'>

const Drawer = createDrawerNavigator<DrawerParams>()
const Stack = createStackNavigator()


export const DrawerNavigatorRoutes = (props: Props) => {
    return (
        <Drawer.Navigator>
            <Drawer.Screen name={"Home"} component={HomeScreen}
                           options={{drawerLabel: 'Home', title: 'Home', headerTitle: 'Home'}}
            />
            <Drawer.Screen name={"AddNewSpend"} component={AddNewSpendScreen}
                           options={{drawerLabel: 'Add New', headerShown: false}}/>
        </Drawer.Navigator>
    )
}

import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {DrawerParams} from "../types";
import {createStackNavigator} from "@react-navigation/stack";
import {HomeStackParams} from "./types";
import {HomeScreen} from "./home";
import {AddNewSpendScreen} from "./add-spending";
import {ViewSpending} from "./edit-spending";


type Props = NativeStackScreenProps<DrawerParams, 'Home'>

const Stack = createStackNavigator<HomeStackParams>()


export const HomeRoutes = (props: Props) => {
    return (
        <Stack.Navigator>
            <Stack.Screen name={'Home'} component={HomeScreen}/>
            <Stack.Screen name={'AddNewSpend'} component={AddNewSpendScreen}/>
            <Stack.Screen name={'ViewSpend'} component={ViewSpending}/>
        </Stack.Navigator>
    )
}



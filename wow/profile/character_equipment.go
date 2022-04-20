package profile

import (
	"encoding/json"
	"fmt"
)

type WowCharacterEquipment struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	EquippedItems []struct {
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"item"`
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		Quantity  int   `json:"quantity"`
		Context   int   `json:"context"`
		BonusList []int `json:"bonus_list,omitempty"`
		Quality   struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Name  string `json:"name"`
		Media struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"media"`
		ItemClass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"item_class"`
		ItemSubclass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"item_subclass"`
		InventoryType struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"inventory_type"`
		Binding struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"binding"`
		Armor struct {
			Value   int `json:"value"`
			Display struct {
				DisplayString string `json:"display_string"`
				Color         struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float64 `json:"a"`
				} `json:"color"`
			} `json:"display"`
		} `json:"armor,omitempty"`
		Stats []struct {
			Type struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"type"`
			Value   int `json:"value"`
			Display struct {
				DisplayString string `json:"display_string"`
				Color         struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float64 `json:"a"`
				} `json:"color"`
			} `json:"display"`
			IsEquipBonus bool `json:"is_equip_bonus,omitempty"`
		} `json:"stats,omitempty"`
		SellPrice struct {
			Value          int `json:"value"`
			DisplayStrings struct {
				Header string `json:"header"`
				Gold   string `json:"gold"`
				Silver string `json:"silver"`
				Copper string `json:"copper"`
			} `json:"display_strings"`
		} `json:"sell_price,omitempty"`
		Requirements struct {
			Level struct {
				Value         int    `json:"value"`
				DisplayString string `json:"display_string"`
			} `json:"level"`
			PlayableClasses struct {
				Links []struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"links"`
				DisplayString string `json:"display_string"`
			} `json:"playable_classes"`
		} `json:"requirements,omitempty"`
		Set struct {
			ItemSet struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"item_set"`
			Items []struct {
				Item struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"item"`
				IsEquipped bool `json:"is_equipped,omitempty"`
			} `json:"items"`
			Effects []struct {
				DisplayString string `json:"display_string"`
				RequiredCount int    `json:"required_count"`
				IsActive      bool   `json:"is_active"`
			} `json:"effects"`
			DisplayString string `json:"display_string"`
		} `json:"set,omitempty"`
		Level struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"level"`
		Transmog struct {
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"item"`
			DisplayString            string `json:"display_string"`
			ItemModifiedAppearanceID int    `json:"item_modified_appearance_id"`
		} `json:"transmog,omitempty"`
		Durability struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"durability,omitempty"`
		Sockets []struct {
			SocketType struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"socket_type"`
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"item"`
			DisplayString string `json:"display_string"`
			Media         struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID int `json:"id"`
			} `json:"media"`
		} `json:"sockets,omitempty"`
		LimitCategory string `json:"limit_category,omitempty"`
		Spells        []struct {
			Spell struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"spell"`
			Description string `json:"description"`
		} `json:"spells,omitempty"`
		Description      string `json:"description,omitempty"`
		IsSubclassHidden bool   `json:"is_subclass_hidden,omitempty"`
		Enchantments     []struct {
			DisplayString string `json:"display_string"`
			SourceItem    struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"source_item"`
			EnchantmentID   int `json:"enchantment_id"`
			EnchantmentSlot struct {
				ID   int    `json:"id"`
				Type string `json:"type"`
			} `json:"enchantment_slot"`
		} `json:"enchantments,omitempty"`
		TimewalkerLevel      int `json:"timewalker_level,omitempty"`
		ModifiedCraftingStat []struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"modified_crafting_stat,omitempty"`
		ModifiedAppearanceID int `json:"modified_appearance_id,omitempty"`
		NameDescription      struct {
			DisplayString string `json:"display_string"`
			Color         struct {
				R int     `json:"r"`
				G int     `json:"g"`
				B int     `json:"b"`
				A float64 `json:"a"`
			} `json:"color"`
		} `json:"name_description,omitempty"`
		UniqueEquipped string `json:"unique_equipped,omitempty"`
		Weapon         struct {
			Damage struct {
				MinValue      int    `json:"min_value"`
				MaxValue      int    `json:"max_value"`
				DisplayString string `json:"display_string"`
				DamageClass   struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"damage_class"`
			} `json:"damage"`
			AttackSpeed struct {
				Value         int    `json:"value"`
				DisplayString string `json:"display_string"`
			} `json:"attack_speed"`
			Dps struct {
				Value         float64 `json:"value"`
				DisplayString string  `json:"display_string"`
			} `json:"dps"`
		} `json:"weapon,omitempty"`
	} `json:"equipped_items"`
	EquippedItemSets []struct {
		ItemSet struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"item_set"`
		Items []struct {
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"item"`
			IsEquipped bool `json:"is_equipped,omitempty"`
		} `json:"items"`
		Effects []struct {
			DisplayString string `json:"display_string"`
			RequiredCount int    `json:"required_count"`
			IsActive      bool   `json:"is_active"`
		} `json:"effects"`
		DisplayString string `json:"display_string"`
	} `json:"equipped_item_sets"`
}

func (c *CharacterClient) Equipment() (WowCharacterEquipment, error) {
	var w WowCharacterEquipment
	body, err := c.Caller("/equipment")
	if err != nil {
		return w, err
	}

	if err = json.Unmarshal(body, &w); err != nil {
		return w, fmt.Errorf("unable to fetch character from returned API values")
	}

	return w, nil
}

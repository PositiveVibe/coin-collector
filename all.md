Below is a more detailed breakdown of each lesson with example code snippets and guidance on **where** to place them in your existing `Scene2` (or a new scene, if you prefer). In many places, you’ll add new variables in the `create()` method, load additional assets in `preload()`, and handle logic in `update()` or inside collision callbacks. Adjust any names or sprite references to match your actual files and style.

---

# Lesson 1: Health Bar for the Player

### Goal
Display a health bar that updates whenever the player takes damage or is healed.

### Steps

1. **Add Health Variables**

   In `create()` (somewhere near where you set up your player), define:
   ```js
   // Player health
   this.maxHealth = 100;
   this.playerHealth = 100;
   ```

2. **Draw the Health Bar**

   We’ll use Phaser’s `Graphics` object to draw a rectangle at a fixed screen position.  
   Still in `create()`:

   ```js
   // Create a health bar background
   this.healthBarBG = this.add.graphics();
   this.healthBarBG.fillStyle(0x000000, 1);  // black
   // x=20, y=20 for example, width=200, height=20
   this.healthBarBG.fillRect(20, 20, 200, 20); 
   this.healthBarBG.setScrollFactor(0); // so it stays in place as the camera moves

   // Create a health bar fill
   this.healthBarFill = this.add.graphics();
   this.healthBarFill.setScrollFactor(0); 
   ```

3. **Create an Update Function for the Health Bar**

   Also in `create()` or at the bottom of your scene, add a function:

   ```js
   this.updateHealthBar = function() {
       // Clear existing fill
       this.healthBarFill.clear();

       // Calculate percentage
       let healthRatio = this.playerHealth / this.maxHealth;
       if (healthRatio < 0) healthRatio = 0;

       // Draw the fill (red)
       this.healthBarFill.fillStyle(0xff0000, 1);
       this.healthBarFill.fillRect(20, 20, 200 * healthRatio, 20);
   };
   ```

4. **Call `this.updateHealthBar()` Whenever Health Changes**

   For example, in your orc collision callback:
   ```js
   function orcPlayerCollision(player, orc) {
     if (this.player.anims.currentAnim && 
         this.player.anims.currentAnim.key === 'attack' && 
         this.player.anims.isPlaying) {
       // We kill the orc
       orc.disableBody(true, true);
     } else {
       // Player takes damage
       this.playerHealth -= 20; // for example
       this.updateHealthBar();  // update the bar

       // Check if player is dead
       if (this.playerHealth <= 0) {
         this.physics.pause();
         this.player.setTint(0xff0000);
         this.player.anims.play('turn');
         this.gameOver = true;
       }
     }
   }
   ```
   Do the same in `hitBomb()` or anywhere the player loses health.

5. **(Optional) Heal the Player**
   If you have healing items later, just add:
   ```js
   this.playerHealth = Math.min(this.playerHealth + 10, this.maxHealth);
   this.updateHealthBar();
   ```

---

# Lesson 2: New Enemy — Dragons (Blasting Lava/Fire)

### Goal
Introduce a dragon enemy that hovers or moves around and shoots fire projectiles at the player.

### Steps

1. **Load Dragon Assets (in `preload()`)**
   ```js
   this.load.spritesheet('dragon', 'assets/dragon.png', {
     frameWidth: 100,
     frameHeight: 100
   });
   this.load.image('dragonFire', 'assets/dragonFire.png');
   ```

2. **Create a Dragon Group (in `create()`)**
   ```js
   this.dragons = this.physics.add.group();
   // Example: spawn one or multiple dragons
   let dragon = this.dragons.create(1200, 300, 'dragon');
   dragon.setScale(2);
   dragon.body.setGravityY(0);   // maybe they fly
   dragon.setCollideWorldBounds(true);
   dragon.health = 5;  // boss-level health or normal

   // Animate the dragon
   this.anims.create({
     key: 'dragonFly',
     frames: this.anims.generateFrameNumbers('dragon', { start: 0, end: 5 }),
     frameRate: 10,
     repeat: -1
   });

   this.dragons.children.iterate((d) => {
     d.anims.play('dragonFly', true);
     d.setVelocityX(-100); // or any movement you want
   });
   ```

3. **Dragon Fire Projectiles (in `create()`)**
   ```js
   this.dragonFireGroup = this.physics.add.group();

   // Periodically spawn fire from each dragon
   this.time.addEvent({
     delay: 2000,             // fire every 2 seconds
     callback: () => {
       this.dragons.children.iterate((dragon) => {
         if (!dragon.active) return; // skip dead dragons

         // Create a fire projectile
         let fire = this.dragonFireGroup.create(dragon.x, dragon.y, 'dragonFire');
         fire.setScale(0.5);
         fire.body.allowGravity = false; 
         // Aim at the player
         this.physics.moveTo(fire, this.player.x, this.player.y, 200);
       });
     },
     loop: true
   });
   ```

4. **Handle Fire Collision with Player**
   ```js
   this.physics.add.overlap(
     this.player,
     this.dragonFireGroup,
     (player, fire) => {
       // Damage player
       this.playerHealth -= 10;
       this.updateHealthBar();
       // Remove the fire projectile
       fire.destroy();
       // Check for death
       if (this.playerHealth <= 0) {
         // gameOver logic
       }
     },
     null,
     this
   );
   ```

5. **(Optional) Collisions with Platforms**  
   If you want the dragons to land or be blocked by platforms:
   ```js
   this.physics.add.collider(this.dragons, this.platforms);
   ```

---

# Lesson 3: New Enemy — Exploding Bombs

### Goal
Create bombs that differ from your star-spawned bombs, e.g. they chase the player or explode on a timer.

### Steps

1. **In `preload()`**
   ```js
   this.load.image('explodingBomb', 'assets/explodingBomb.png');
   this.load.spritesheet('explosion', 'assets/explosion.png', {
     frameWidth: 64,
     frameHeight: 64
   });
   ```

2. **In `create()` — Setup Exploding Bombs Group**
   ```js
   this.explodingBombs = this.physics.add.group();
   // Example: spawn one bomb
   let eBomb = this.explodingBombs.create(300, 100, 'explodingBomb');
   eBomb.body.setBounce(0.5);
   eBomb.setCollideWorldBounds(true);

   // Explosion animation
   this.anims.create({
     key: 'explodeAnim',
     frames: this.anims.generateFrameNumbers('explosion', { start: 0, end: 5 }),
     frameRate: 10,
     repeat: 0
   });

   // If you want the bomb to chase the player
   // (You could also do this in `update()`)
   this.time.addEvent({
     delay: 100,
     callback: () => {
       this.explodingBombs.children.iterate((bomb) => {
         this.physics.moveToObject(bomb, this.player, 100);
       });
     },
     loop: true
   });

   // Collide with platforms
   this.physics.add.collider(this.explodingBombs, this.platforms);

   // Overlap with player -> explosion
   this.physics.add.overlap(
     this.player,
     this.explodingBombs,
     this.handleBombExplode,
     null,
     this
   );
   ```

3. **Explosion Callback**
   ```js
   this.handleBombExplode = function(player, bomb) {
     // Temporarily disable bomb physics
     bomb.body.enable = false;
     bomb.setVelocity(0);

     // Show explosion animation
     let explosionSprite = this.add.sprite(bomb.x, bomb.y, 'explosion');
     explosionSprite.play('explodeAnim');

     // Destroy bomb after animation finishes
     explosionSprite.on('animationcomplete', () => {
       explosionSprite.destroy();
     });
     bomb.destroy();

     // Damage player
     this.playerHealth -= 30;
     this.updateHealthBar();
     if (this.playerHealth <= 0) {
       // handle game over
     }
   };
   ```

---

# Lesson 4: Boss Monster — Ice Monster (Yeti)

### Goal
A “boss” character with multiple attacks: ice, wind, and summoning Angry Snowmen.

### Steps

1. **Load Assets in `preload()`**
   ```js
   this.load.spritesheet('yetiIdle', 'assets/yetiIdle.png', { frameWidth: 100, frameHeight: 100 });
   this.load.spritesheet('yetiAttack', 'assets/yetiAttack.png', { frameWidth: 100, frameHeight: 100 });
   this.load.image('iceProjectile', 'assets/iceProjectile.png');
   this.load.spritesheet('angrySnowman', 'assets/angrySnowman.png', { frameWidth: 50, frameHeight: 50 });
   ```

2. **Create the Yeti (in `create()`)**
   ```js
   this.yeti = this.physics.add.sprite(1400, 1000, 'yetiIdle');
   this.yeti.health = 20;
   this.yeti.setScale(2);
   this.yeti.setCollideWorldBounds(true);

   // Animations (idle, attack)
   this.anims.create({
     key: 'yetiIdleAnim',
     frames: this.anims.generateFrameNumbers('yetiIdle', { start: 0, end: 5 }),
     frameRate: 5,
     repeat: -1
   });
   this.anims.create({
     key: 'yetiAttackAnim',
     frames: this.anims.generateFrameNumbers('yetiAttack', { start: 0, end: 5 }),
     frameRate: 10,
     repeat: 0
   });

   this.yeti.anims.play('yetiIdleAnim', true);

   // Yeti Attack Timer
   this.time.addEvent({
     delay: 3000,
     callback: this.yetiAttackPattern,
     loop: true,
     callbackScope: this
   });

   // Collisions
   this.physics.add.collider(this.yeti, this.platforms);
   this.physics.add.overlap(this.player, this.yeti, this.onPlayerHitByYeti, null, this);
   ```

3. **Define the Attack Pattern**  
   We can randomly choose an attack: shoot ice, create wind, or summon snowmen.

   ```js
   this.yetiAttackPattern = function() {
     if (!this.yeti.active) return;
     let attackType = Phaser.Math.Between(1, 3); // 1 = ice, 2 = wind, 3 = snowmen

     switch(attackType) {
       case 1: 
         this.shootIce();
         break;
       case 2:
         this.castWind();
         break;
       case 3:
         this.summonSnowmen();
         break;
     }
   };
   ```

4. **Implement Each Attack**

   ```js
   // 1) Shoot Ice
   this.shootIce = function() {
     this.yeti.anims.play('yetiAttackAnim');
     let ice = this.physics.add.sprite(this.yeti.x, this.yeti.y, 'iceProjectile');
     this.physics.moveToObject(ice, this.player, 200);
     // Overlap to damage player
     this.physics.add.overlap(this.player, ice, () => {
       this.playerHealth -= 10;
       this.updateHealthBar();
       ice.destroy();
     }, null, this);
   };

   // 2) Cast Wind (push the player away)
   this.castWind = function() {
     this.yeti.anims.play('yetiAttackAnim');
     // Quick example: set player velocity away from the Yeti
     let direction = Math.sign(this.player.x - this.yeti.x);
     this.player.setVelocityX(600 * direction);
   };

   // 3) Summon Snowmen
   this.summonSnowmen = function() {
     this.yeti.anims.play('yetiAttackAnim');
     this.angrySnowmen = this.angrySnowmen || this.physics.add.group();
     let snowman = this.angrySnowmen.create(this.yeti.x, this.yeti.y, 'angrySnowman');
     snowman.setCollideWorldBounds(true);
     snowman.setVelocityX(Phaser.Math.Between(-100, 100));
     snowman.health = 3;
     // Add collision logic
     this.physics.add.collider(snowman, this.platforms);
     this.physics.add.overlap(this.player, snowman, (player, sn) => {
       this.playerHealth -= 5;
       this.updateHealthBar();
       if (this.playerHealth <= 0) {
         // game over
       }
     }, null, this);
   };
   ```

5. **Player vs. Yeti Collision**
   ```js
   this.onPlayerHitByYeti = function(player, yeti) {
     // If attacking, damage Yeti
     if (player.anims.currentAnim && player.anims.currentAnim.key === 'attack') {
       yeti.health--;
       if (yeti.health <= 0) {
         yeti.disableBody(true, true);
         // maybe remove from scene or show defeat animation
       }
     } else {
       // Player takes damage
       this.playerHealth -= 10;
       this.updateHealthBar();
       if (this.playerHealth <= 0) {
         // game over
       }
     }
   };
   ```

---

# Lesson 5: Actual Acid Dragon

### Goal
A variant of the dragon with “acid asteroids” and “acid spit” attacks.

### Steps

1. **In `preload()`**
   ```js
   this.load.spritesheet('acidDragon', 'assets/acidDragon.png', { frameWidth: 100, frameHeight: 100 });
   this.load.image('acidAsteroid', 'assets/acidAsteroid.png');
   this.load.image('acidSpit', 'assets/acidSpit.png');
   ```

2. **Create the Acid Dragon in `create()`**
   ```js
   this.acidDragon = this.physics.add.sprite(1000, 200, 'acidDragon');
   this.acidDragon.health = 10;
   this.acidDragon.setScale(2);
   this.acidDragon.setCollideWorldBounds(true);

   this.anims.create({
     key: 'acidDragonFly',
     frames: this.anims.generateFrameNumbers('acidDragon', { start: 0, end: 5 }),
     frameRate: 10,
     repeat: -1
   });
   this.acidDragon.play('acidDragonFly');

   // Timed attack
   this.time.addEvent({
     delay: 2500,
     callback: this.acidDragonAttack,
     loop: true,
     callbackScope: this
   });
   ```

3. **Attack Patterns (in `create()` or after it)**
   ```js
   this.acidDragonAttack = function() {
     if (!this.acidDragon.active) return;

     // 50% chance acid asteroid, 50% acid spit
     if (Math.random() < 0.5) {
       this.shootAcidAsteroid();
     } else {
       this.shootAcidSpit();
     }
   };

   this.shootAcidAsteroid = function() {
     let asteroid = this.physics.add.sprite(this.acidDragon.x, this.acidDragon.y, 'acidAsteroid');
     asteroid.setScale(0.7);
     asteroid.body.allowGravity = false;
     // Perhaps a downward arc
     asteroid.setVelocity(Phaser.Math.Between(-100, 100), 200);

     this.physics.add.overlap(this.player, asteroid, () => {
       // Big damage
       this.playerHealth -= 20;
       this.updateHealthBar();
       asteroid.destroy();
       // If you want, do area damage or an explosion
     }, null, this);
   };

   this.shootAcidSpit = function() {
     let spit = this.physics.add.sprite(this.acidDragon.x, this.acidDragon.y, 'acidSpit');
     spit.body.allowGravity = false;
     // Possibly one-shot kill if you want
     this.physics.moveToObject(spit, this.player, 300);

     this.physics.add.overlap(this.player, spit, () => {
       this.playerHealth = 0; // one-shot example
       this.updateHealthBar();
       spit.destroy();
       // check gameOver
     }, null, this);
   };
   ```

4. **Player Attacking the Acid Dragon**
   ```js
   this.physics.add.overlap(this.player, this.acidDragon, (player, dragon) => {
     if (player.anims.currentAnim && player.anims.currentAnim.key === 'attack') {
       dragon.health--;
       if (dragon.health <= 0) {
         dragon.disableBody(true, true);
         // defeat
       }
     } else {
       // Player is damaged by contact
       this.playerHealth -= 10;
       this.updateHealthBar();
       // check if <= 0
     }
   }, null, this);
   ```

---

# Lesson 6: Power-Ups — Mystery Potion (Bad Effects)

### Goal
When the player collects a “mystery potion,” they get a random negative effect.  

**Bad Effects**:
1. Makes clone helper attack you  
2. Slows you down  
3. Decreases health  
4. Turns off gravity  

### Steps

1. **In `preload()`**
   ```js
   this.load.image('mysteryPotion', 'assets/mysteryPotion.png');
   ```

2. **In `create()` — Mystery Potion Group**
   ```js
   this.mysteryPotions = this.physics.add.group({
     key: 'mysteryPotion',
     repeat: 3,
     setXY: { x: 300, y: 100, stepX: 400 }
   });
   this.physics.add.collider(this.mysteryPotions, this.platforms);

   // Overlap with player
   this.physics.add.overlap(
     this.player,
     this.mysteryPotions,
     this.handleMysteryPotion,
     null,
     this
   );
   ```

3. **Handle Mystery Potion**
   ```js
   this.handleMysteryPotion = function(player, potion) {
     potion.disableBody(true, true);

     // List of negative effects
     const badEffects = ['cloneAttack', 'slow', 'healthDown', 'noGravity'];
     let chosen = Phaser.Utils.Array.GetRandom(badEffects);

     switch(chosen) {
       case 'cloneAttack':
         this.turnCloneHelperEvil();
         break;
       case 'slow':
         this.applySlowEffect();
         break;
       case 'healthDown':
         this.playerHealth -= 20;
         this.updateHealthBar();
         break;
       case 'noGravity':
         this.player.body.allowGravity = false;
         // Re-enable after 5 seconds
         this.time.delayedCall(5000, () => {
           this.player.body.allowGravity = true;
         });
         break;
     }
   };
   ```

4. **Implement Each Bad Effect**  
   ```js
   this.turnCloneHelperEvil = function() {
     // If you have a clone sprite that normally helps you,
     // set a flag so it attacks the player.
     // Example:
     if (this.cloneHelper) {
       this.cloneHelper.target = this.player; 
       this.cloneHelper.isEvil = true;
     }
   };

   this.applySlowEffect = function() {
     // Store original speed somewhere, e.g. this.playerSpeed
     let originalSpeed = 500;
     // Reduce speed
     this.playerSlowSpeed = 200;
     // For 5 seconds
     this.time.delayedCall(5000, () => {
       this.playerSlowSpeed = originalSpeed;
     });
   };
   ```

5. **Adjust Movement in `update()`**  
   If you’re controlling the player via `player.setVelocityX(500)`, you can reference `this.playerSlowSpeed`:
   ```js
   if (cursors.left.isDown) {
     player.setVelocityX(-this.playerSlowSpeed || -500);
     // ...
   }
   ```
   That way, if `applySlowEffect()` changes `playerSlowSpeed`, it affects the player.

---

# Lesson 7: Power-Ups — Mystery Potion (Good Effects)

### Goal
A similar “mystery potion” system, but with random **positive** effects.

**Good Effects**:
1. Speed  
2. Jump Boost  
3. Clone Helper  
4. Healing Potion  
5. Coin Magnet  

### Steps

1. **In `create()` — Good Mystery Potions**  
   You can combine them with the same `mysteryPotion` sprite or use a separate sprite. Maybe do a `50%` chance it’s good vs. bad:

   ```js
   this.physics.add.overlap(
     this.player,
     this.mysteryPotions,
     this.handleMysteryPotion, 
     null, 
     this
   );

   // Inside the same handleMysteryPotion:
   this.handleMysteryPotion = function(player, potion) {
     potion.disableBody(true, true);

     let isGood = Math.random() < 0.5; // 50% chance

     if (isGood) {
       let goodEffects = ['speed', 'jump', 'cloneHelper', 'heal', 'coinMagnet'];
       let chosenGood = Phaser.Utils.Array.GetRandom(goodEffects);
       this.applyGoodEffect(chosenGood);
     } else {
       let badEffects = ['cloneAttack','slow','healthDown','noGravity'];
       let chosenBad = Phaser.Utils.Array.GetRandom(badEffects);
       this.applyBadEffect(chosenBad);
     }
   };
   ```

2. **Apply Good Effects**
   ```js
   this.applyGoodEffect = function(effect) {
     switch(effect) {
       case 'speed':
         this.applySpeedBoost();
         break;
       case 'jump':
         this.applyJumpBoost();
         break;
       case 'cloneHelper':
         this.spawnCloneHelper();
         break;
       case 'heal':
         this.playerHealth = Math.min(this.playerHealth + 30, this.maxHealth);
         this.updateHealthBar();
         break;
       case 'coinMagnet':
         this.enableCoinMagnet();
         break;
     }
   };
   ```

3. **Implement Each Good Effect**

   - **Speed Boost**:
     ```js
     this.applySpeedBoost = function() {
       // Increase movement speed to, say, 800
       this.playerSpeed = 800;
       // revert after 5 seconds
       this.time.delayedCall(5000, () => {
         this.playerSpeed = 500; 
       });
     };
     ```
     And in your `update()` logic, use `this.playerSpeed` for setting velocityX.

   - **Jump Boost**:
     ```js
     this.applyJumpBoost = function() {
       this.jumpHeight = -900; // originally -600
       this.time.delayedCall(5000, () => {
         this.jumpHeight = -600;
       });
     };
     ```
     Then in `if (cursors.up.isDown && player.body.touching.down)` do:
     ```js
     player.setVelocityY(this.jumpHeight || -600);
     ```

   - **Clone Helper**:
     ```js
     this.spawnCloneHelper = function() {
       this.cloneHelper = this.physics.add.sprite(this.player.x + 50, this.player.y, 'dude');
       this.cloneHelper.setScale(2);
       // Simple AI: follow the player or move around
       this.physics.add.overlap(this.cloneHelper, this.orcs, (clone, orc) => {
         // Attack the orc
         orc.disableBody(true, true);
       }, null, this);
     };
     ```

   - **Healing Potion**: (already covered above, just add to `this.playerHealth`)

   - **Coin Magnet**:
     ```js
     this.enableCoinMagnet = function() {
       // Suppose your coins/stars are in this.stars
       this.magnetActive = true;
       this.time.delayedCall(5000, () => { 
         this.magnetActive = false; 
       });
     };

     // Then in update():
     if (this.magnetActive) {
       this.stars.children.iterate((star) => {
         // Pull each star toward player
         this.physics.moveToObject(star, this.player, 100);
       });
     }
     ```

---

# Putting It All Together

1. **Scene Organization**  
   - Most code is placed in `create()` for setup: loading groups, collisions, timers, etc.  
   - Ongoing AI or effect logic goes in `update()` or inside timed events.  
   - Collision callbacks are often defined inline or as separate functions.

2. **Variable Management**  
   - Consider storing `this.playerSpeed = 500; this.jumpHeight = -600;` so they can be temporarily changed by power-ups.

3. **Testing & Tweaking**  
   - Encourage the student to adjust speeds, damage, spawn frequencies, and effect durations to balance the game.

With these code snippets, you can copy-paste them into your `Scene2` (or additional scenes) at the indicated sections (`preload()`, `create()`, `update()`, or separate helper functions). This structure will guide the student through each feature step-by-step, showing exactly how and where to implement new functionality.
